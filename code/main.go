package main

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// Run the collector in a goroutine
	go func() {
		if err := col.Run(ctx); err != nil {
			logger.Error("Collector exited with error", zap.Error(err))
		}
	}()

	// Simulate running for some time
	time.Sleep(5 * time.Second)

	// Stop the collector gracefully
	col.Stop()

	// Wait for shutdown to complete
	time.Sleep(1 * time.Second)
	logger.Info("Main function exiting")
}

// Run starts the collector and listens for shutdown signals.
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

// Stop gracefully shuts down the collector by sending a signal to shutdownChan.
func (col *Collector) Stop() {
	col.logger.Info("Stopping collector...")
	close(col.shutdownChan) // Send shutdown signal
}

// shutdown performs cleanup operations before the collector exits.
func (col *Collector) shutdown(ctx context.Context) error {
	col.logger.Info("Shutting down collector")
	// Add cleanup logic here, e.g., closing files, releasing resources, etc.
	return nil
}
