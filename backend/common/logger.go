package common

import (
	"log/slog"
	"os"
	"sync"
)

var (
	logger  *slog.Logger
	onceLog sync.Once
)

func LoadLogger() {
	cfg := GetConfig()

	var logHandler slog.Handler
	var logOpts slog.HandlerOptions

	switch cfg.ENV {
	case EnvDevelopment:
		logOpts.Level = slog.LevelDebug
		logOpts.AddSource = true
		logHandler = slog.NewTextHandler(os.Stdout, &logOpts)
	case EnvProduction:
		logOpts.Level = slog.LevelInfo
		logHandler = slog.NewJSONHandler(os.Stdout, &logOpts)
	default:
		logOpts.Level = slog.LevelInfo
		logHandler = slog.NewJSONHandler(os.Stdout, &logOpts)
	}

	log := slog.New(logHandler)
	slog.SetDefault(log)
	onceLog.Do(func() {
		logger = log
	})
}

func GetLogger() *slog.Logger {
	if logger == nil {
		panic("Global logger not initialized. Call LoadLogger() first.")
	}
	return logger
}
