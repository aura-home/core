package logger

import (
	"fmt"
	"github.com/aura-home/core/internal/config"
	"log/slog"

	"os"
	"strings"
)

type Logger struct {
	log *slog.Logger
	cfg *config.LoggerConfig
}

func New(cfg *config.LoggerConfig) *Logger {
	var handler slog.Handler

	if cfg.Type == "json" {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLevel(cfg.Level),
		})
	}

	logger := slog.New(handler)

	return &Logger{
		cfg: cfg,
		log: logger,
	}
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	if l.cfg.Enable {
		l.log.Info(msg, keysAndValues...)
	}
}

func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	if l.cfg.Enable {
		l.log.Error(fmt.Sprintf("%s: %v", msg, err), keysAndValues...)
	}
}

func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	if l.cfg.Enable {
		l.log.Debug(msg, keysAndValues...)
	}
}

func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	if l.cfg.Enable {
		l.log.Warn(msg, keysAndValues...)
	}
}
