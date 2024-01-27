package main

import (
	"fmt"

	"github.com/banderveloper/go-forms/internal/config"
	"github.com/banderveloper/go-forms/internal/lib/jwthandler"
)

func main() {

	// Initialize configuration
	cfg := config.MustLoad()

	// Initialize pretty logger based on slog
	// slogger := logger.New(cfg.Environment)

	jwtHandler := jwthandler.New(cfg)

	fmt.Println(jwtHandler.GetAccessToken(55))
	fmt.Println(jwtHandler.GetRefreshToken(56))

	// fmt.Println(cfg)
}
