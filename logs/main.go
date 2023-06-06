package main

import (
	"os"
	"time"

	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type yhconfiglog struct {
	Level      string
	Time       time.Time
	Message    string
	Stacktrace string
}

type PgHook struct {
	db *pg.DB
}

func (h *PgHook) Write(p []byte) (int, error) {
	logEntry := yhconfiglog{
		Level:      "debug",
		Time:       time.Now(),
		Message:    string(p),
		Stacktrace: "",
	}
	_, err := h.db.Model(&logEntry).Insert()
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (h *PgHook) Sync() error {
	return nil
}

func main() {
	// Connect to the database.
	db := pg.Connect(&pg.Options{})
	defer db.Close()

	// Create a Zap logger with a custom encoder and a core with a Zap Hook.
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	pgHook := &PgHook{db: db}
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(pgHook))

	core := zapcore.NewCore(
		encoder,
		multiWriteSyncer,
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	logger := zap.New(core)
	defer logger.Sync()

	// yhconfiglog some messages.
	logger.Warn("A warning occurred.")
	logger.Info("Some information.")
	logger.Error("An error occurred.")

	// Wait a bit for the log message to be written to the database.
	time.Sleep(2 * time.Second)

	// Query the logs from the database.
	//var logs []yhconfiglog
	//err := db.Model(&logs).Select()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Print the logs to the console.
	//for _, log := range logs {
	//	fmt.Printf("%v: [%v] %v\n%s\n", log.Time, log.Level, log.Message, log.Stacktrace)
	//}
}
