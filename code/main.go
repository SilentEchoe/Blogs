package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Worker 结构体，模拟一个工作器
type Worker struct {
	signalsChannel      chan os.Signal
	asyncErrorChannel   chan error
	reloadConfigChannel chan struct{}
	shutdownChan        chan struct{}
	service             struct {
		Logger func() *zap.Logger
	}
	// 模拟配置
	config struct {
		// 工作间隔
		Interval time.Duration
	}
}

// NewWorker 创建一个新的 Worker 实例
func NewWorker() *Worker {
	logger, _ := zap.NewProduction()
	w := &Worker{
		signalsChannel:      make(chan os.Signal, 1),
		asyncErrorChannel:   make(chan error, 1),
		reloadConfigChannel: make(chan struct{}),
		shutdownChan:        make(chan struct{}),
	}
	w.service.Logger = func() *zap.Logger {
		return logger
	}
	w.config.Interval = 2 * time.Second
	return w
}

// setupConfigurationComponents 模拟设置配置组件
func (w *Worker) setupConfigurationComponents(ctx context.Context) error {
	// 模拟一些初始化工作
	w.service.Logger().Info("Setting up configuration components")
	time.Sleep(1 * time.Second)
	return nil
}

// reloadConfiguration 模拟重新加载配置
func (w *Worker) reloadConfiguration(ctx context.Context) error {
	w.service.Logger().Info("Reloading configuration")
	// 模拟重新加载配置的耗时操作
	time.Sleep(1 * time.Second)
	// 这里可以更新配置
	w.config.Interval = 3 * time.Second
	return nil
}

// shutdown 模拟关闭操作
func (w *Worker) shutdown(ctx context.Context) error {
	w.service.Logger().Info("Shutting down worker")
	// 模拟关闭操作的耗时
	time.Sleep(1 * time.Second)
	return nil
}

// Run 启动工作器并等待其完成
func (w *Worker) Run(ctx context.Context) error {
	// 初始化配置组件
	if err := w.setupConfigurationComponents(ctx); err != nil {
		return err
	}

	// 始终监听 SIGHUP 信号以进行配置重新加载
	signal.Notify(w.signalsChannel, syscall.SIGHUP)
	defer signal.Stop(w.signalsChannel)

	// 监听 SIGTERM 和 SIGINT 信号以进行优雅关闭
	signal.Notify(w.signalsChannel, os.Interrupt, syscall.SIGTERM)

	// 启动一个 goroutine 模拟异步工作
	go func() {
		ticker := time.NewTicker(w.config.Interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				w.service.Logger().Info("Doing some work")
				// 模拟异步错误
				if time.Now().Unix()%5 == 0 {
					w.asyncErrorChannel <- fmt.Errorf("simulated async error")
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// 控制循环：监听各种中断信号
LOOP:
	for {
		select {
		case err := <-w.asyncErrorChannel:
			w.service.Logger().Error("Asynchronous error received, terminating process", zap.Error(err))
			break LOOP
		case s := <-w.signalsChannel:
			w.service.Logger().Info("Received signal from OS", zap.String("signal", s.String()))
			if s == syscall.SIGHUP {
				if err := w.reloadConfiguration(ctx); err != nil {
					return err
				}
			} else {
				break LOOP
			}
		case <-w.shutdownChan:
			w.service.Logger().Info("Received shutdown request")
			break LOOP
		case <-ctx.Done():
			w.service.Logger().Info("Context done, terminating process", zap.Error(ctx.Err()))
			// 使用背景上下文调用关闭函数，因为传入的上下文已被取消
			return w.shutdown(context.Background())
		}
	}

	return w.shutdown(ctx)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	worker := NewWorker()
	if err := worker.Run(ctx); err != nil {
		fmt.Printf("Error running worker: %v\n", err)
	}
}
