package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Collector struct {
	// shutdownChan is used to terminate the collector.
	shutdownChan chan struct{}
	// signalsChannel is used to receive termination signals from the OS.
	signalsChannel chan os.Signal
	logger         *zap.Logger
}

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		return
	}
	defer logger.Sync()

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new collector
	col := &Collector{
		shutdownChan:   make(chan struct{}),
		signalsChannel: make(chan os.Signal, 1),
		logger:         logger,
	}

	// Run the collector
	if err := col.Run(ctx); err != nil {
		logger.Error("Collector exited with error", zap.Error(err))
	}
}

func (col *Collector) Run(ctx context.Context) error {
	// Notify signalsChannel for SIGHUP, SIGINT, and SIGTERM
	signal.Notify(col.signalsChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(col.signalsChannel)

	col.logger.Info("Collector started")

	// Control loop: selects between channels for various interrupts
LOOP:
	for {
		select {
		case <-col.shutdownChan:
			col.logger.Info("Received shutdown request")
			break LOOP
		case sig := <-col.signalsChannel:
			col.logger.Info("Received signal", zap.String("signal", sig.String()))
			break LOOP
		case <-ctx.Done():
			col.logger.Info("Context done, terminating process", zap.Error(ctx.Err()))
			break LOOP
		}
	}

	// Perform shutdown
	return col.shutdown(ctx)
}

func (col *Collector) shutdown(ctx context.Context) error {
	col.logger.Info("Shutting down collector")
	// Add cleanup logic here, e.g., closing files, releasing resources, etc.
	return nil
}
