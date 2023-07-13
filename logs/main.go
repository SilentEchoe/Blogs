package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type yhlog struct {
	Level      string
	Time       time.Time
	Message    string
	Stacktrace string
}

type PgHook struct {
	db *gorm.DB
}

func (h *PgHook) Write(p []byte) (int, error) {
	logEntry := yhlog{
		Level:      "debug",
		Time:       time.Now(),
		Message:    string(p),
		Stacktrace: "",
	}
	_, err := h.db.Table("").Model(&logEntry).Create(&logEntry).Rows()
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (h *PgHook) Sync() error {
	return nil
}

var gormOnce = sync.Once{}
var db *gorm.DB

func initGorm() *gorm.DB {
	gormOnce.Do(func() {
		var err error
		dsn := ""
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			fmt.Println("connect postgres error:%v", err)
		}
	})
	return db
}

func main() {
	// Connect to the database.
	initGorm()

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

	// yhlog some messages.
	logger.Warn("A warning occurred.")
	logger.Info("Some information.")
	logger.Error("An error occurred.")

	// Wait a bit for the log message to be written to the database.
	time.Sleep(2 * time.Second)

}
