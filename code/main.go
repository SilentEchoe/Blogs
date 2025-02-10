package main

import (
	"context"
	"fmt"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	mMetric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	googleRpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type MetricType int32

const (
	MetricType_COUNTER         MetricType = 0
	MetricType_GAUGE           MetricType = 1
	MetricType_SUMMARY         MetricType = 2
	MetricType_UNTYPED         MetricType = 3
	MetricType_HISTOGRAM       MetricType = 4
	MetricType_GAUGE_HISTOGRAM MetricType = 5
	Prometheus                            = "prometheus"
)

func main() {
	// 初始化信号通道
	signalsChannel := make(chan os.Signal, 1)
	// 注册要监听的信号
	signal.Notify(signalsChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	var ctx context.Context

	// 初始化连接
	shutdownMeterProvider, err := initMeterProvider(ctx, "0.0.0.0:4317")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	defer func() {
		if err := shutdownMeterProvider(ctx); err != nil {
			fmt.Println("failed to shutdown MeterProvider", zap.Error(err))
		}
	}()

	for i := 0; i < 100; i++ {
		metricFamilies, err := fetchPrometheusMetrics("https://prometheus.demo.do.prometheus.io/metrics")
		if err != nil {
			fmt.Println(fmt.Sprintf("request url err: %v", err))
			continue
		}
		err = recordMetrics(ctx, metricFamilies)
		if err != nil {
			fmt.Println(fmt.Sprintf("record metrics err: %v", err))
			continue
		}
		fmt.Println("消费完毕")
	}

	// 监听信号通道
	sig := <-signalsChannel
	fmt.Printf("接收到信号: %v，程序即将退出...\n", sig)
}

func initMeterProvider(ctx context.Context, metricExportAddress string) (func(context.Context) error, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			semconv.ServiceNameKey.String("test"),
		),
	)
	conn, err := googleRpc.NewClient(metricExportAddress,
		googleRpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(
			metricExporter,
			sdkmetric.WithInterval(time.Duration(10)*time.Second),
		)),
		sdkmetric.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)
	return meterProvider.Shutdown, nil
}

func fetchPrometheusMetrics(url string) (map[string]*io_prometheus_client.MetricFamily, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metrics: %w", err)
	}
	defer resp.Body.Close()

	// 解析 Prometheus 文本格式
	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse metrics: %w", err)
	}

	return metricFamilies, nil
}

func recordMetrics(ctx context.Context, metricFamilies map[string]*io_prometheus_client.MetricFamily) error {
	meter := otel.Meter("edge_exporter")
	for name, family := range metricFamilies {
		for _, metric := range family.Metric {
			// 提取标签
			attrs := make([]attribute.KeyValue, 0, len(metric.Label))
			for _, label := range metric.Label {
				attrs = append(attrs, attribute.String(*label.Name, *label.Value))
			}

			// 根据指标类型记录数据
			switch family.GetType() {
			case io_prometheus_client.MetricType(MetricType_COUNTER):
				counter, err := meter.Float64Counter(name)
				if err != nil {
					return fmt.Errorf("failed to create counter: %w", err)
				}
				counter.Add(ctx, metric.Counter.GetValue(), mMetric.WithAttributes(attrs...))
			case io_prometheus_client.MetricType(MetricType_GAUGE):
				gauge, err := meter.Float64UpDownCounter(name)
				if err != nil {
					return fmt.Errorf("failed to create gauge: %w", err)
				}
				gauge.Add(context.Background(), metric.Gauge.GetValue(), mMetric.WithAttributes(attrs...))
			case io_prometheus_client.MetricType(MetricType_HISTOGRAM):
				histogram, err := meter.Float64Histogram(name + "_histogram")
				if err != nil {
					return fmt.Errorf("failed to create histogram: %w", err)
				}
				for _, bucket := range metric.Histogram.Bucket {
					histogram.Record(context.Background(), bucket.GetUpperBound(), mMetric.WithAttributes(attrs...))
				}
			case io_prometheus_client.MetricType(MetricType_SUMMARY):
				// Prometheus 的 SUMMARY 类型用于记录分布摘要（如分位数）
				summary, err := meter.Float64Histogram(name + "_summary")
				if err != nil {
					return fmt.Errorf("failed to create summary: %w", err)
				}
				// 记录分位数
				for _, quantile := range metric.Summary.Quantile {
					summary.Record(
						context.Background(),
						quantile.GetValue(),
						mMetric.WithAttributes(
							append(attrs, attribute.Float64("quantile", quantile.GetQuantile()))...,
						),
					)
				}
				// 记录总数和总和
				counter, err := meter.Float64Counter(name + "_count")
				if err != nil {
					return fmt.Errorf("failed to create summary count: %w", err)
				}
				counter.Add(
					context.Background(),
					float64(metric.Summary.GetSampleCount()),
					mMetric.WithAttributes(attrs...),
				)
				sum, err := meter.Float64Counter(name + "_sum")
				if err != nil {
					return fmt.Errorf("failed to create summary sum: %w", err)
				}
				sum.Add(
					context.Background(),
					metric.Summary.GetSampleSum(),
					mMetric.WithAttributes(attrs...),
				)

			case io_prometheus_client.MetricType(MetricType_UNTYPED):
				//Prometheus 的 UNTYPED 类型表示未明确类型的指标。通常可以将其视为 GAUGE 类型。
				gauge, err := meter.Float64UpDownCounter(name)
				if err != nil {
					return fmt.Errorf("failed to create untyped gauge: %w", err)
				}
				gauge.Add(
					context.Background(),
					metric.Untyped.GetValue(),
					mMetric.WithAttributes(attrs...),
				)

			case io_prometheus_client.MetricType(MetricType_GAUGE_HISTOGRAM):
				histogram, err := meter.Float64Histogram(name)
				if err != nil {
					return fmt.Errorf("failed to create gauge histogram: %w", err)
				}

				// 记录每个桶的值
				for _, bucket := range metric.Histogram.Bucket {
					histogram.Record(
						context.Background(),
						bucket.GetUpperBound(),
						mMetric.WithAttributes(
							append(attrs, attribute.Float64("le", bucket.GetUpperBound()))...,
						),
					)
				}

				// 记录总数和总和
				counter, err := meter.Float64Counter(name + "_count")
				if err != nil {
					return fmt.Errorf("failed to create gauge histogram count: %w", err)
				}
				counter.Add(
					context.Background(),
					float64(metric.Histogram.GetSampleCount()),
					mMetric.WithAttributes(attrs...),
				)

				sum, err := meter.Float64Counter(name + "_sum")
				if err != nil {
					return fmt.Errorf("failed to create gauge histogram sum: %w", err)
				}
				sum.Add(
					context.Background(),
					metric.Histogram.GetSampleSum(),
					mMetric.WithAttributes(attrs...),
				)
			}

		}

	}
	return nil
}
