package main

import (
	"fmt"

	"github.com/banderveloper/go-forms/internal/config"
	"github.com/banderveloper/go-forms/internal/lib/jwt"
)

func main() {

	// Initialize configuration
	cfg := config.MustLoad()

	// Initialize pretty logger based on slog
	// slogger := logger.New(cfg.Environment)

	jwtHandler := jwt.New(cfg)

	fmt.Println(jwtHandler)

	// fmt.Println(cfg)
}
