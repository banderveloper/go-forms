package main

import (
	"github.com/banderveloper/go-forms/internal/config"
	"github.com/banderveloper/go-forms/internal/lib/jwthandler"
	"github.com/banderveloper/go-forms/internal/lib/logger"
)

func main() {

	// Initialize configuration
	cfg := config.MustLoad()

	// Initialize pretty logger based on slog
	slogger := logger.New(cfg.Environment)
	jwtHandler := jwthandler.New(cfg)

	_, _ = slogger, jwtHandler
}
