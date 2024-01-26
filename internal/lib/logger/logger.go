package logger

import (
	"log/slog"
	"os"
)

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "development"
	ENV_PROD  = "production"
)

// Constructor of logger with integrated pretty printing
func New(env string) *slog.Logger {

	var debugLevel slog.Level

	switch env {
	case ENV_LOCAL:
		debugLevel = slog.LevelDebug
	case ENV_DEV:
	case ENV_PROD:
		debugLevel = slog.LevelDebug
	}

	options := PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: debugLevel,
		},
	}

	handler := options.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func Error(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
