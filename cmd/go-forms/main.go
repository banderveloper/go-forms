package main

import (
	"errors"

	"github.com/banderveloper/go-forms/internal/config"
	"github.com/banderveloper/go-forms/internal/lib/logger"
)

func main() {

	// Initialize configuration
	cfg := config.MustLoad()

	// Initialize pretty logger based on slog
	slogger := logger.New(cfg.Environment)

	// Test error log
	err := errors.ErrUnsupported
	slogger.Error("Logger test", logger.Error(err))

	// fmt.Println(cfg)
}
