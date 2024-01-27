package logger

import (
	"log/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvDev   = "development"
	EnvProd  = "production"
)

// New Constructor of logger with integrated pretty printing
func New(env string) *slog.Logger {

	var debugLevel slog.Level

	switch env {
	case EnvLocal:
		debugLevel = slog.LevelDebug
	case EnvDev:
	case EnvProd:
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
