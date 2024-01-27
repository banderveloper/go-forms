package main

import (
	"github.com/banderveloper/go-forms/internal/config"
	"github.com/banderveloper/go-forms/internal/lib/jwthandler"
	"github.com/banderveloper/go-forms/internal/lib/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"time"
)

func main() {

	// Initialize configuration
	cfg := config.MustLoad()

	// Initialize pretty logger based on slog
	slogger := logger.New(cfg.Environment)
	jwtHandler := jwthandler.New(cfg)

	// Initialize router with middleware and routes
	router := chi.NewRouter()
	setupMiddleware(router)
	setupRoutes(router)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	slogger.Info("Server started", slog.String("address", cfg.Address))

	if err := server.ListenAndServe(); err != nil {
		slogger.Error("failed to start server")
	}

	slogger.Error("server stopped")

	_ = jwtHandler
}

func setupMiddleware(router *chi.Mux) {
	// add requestID for each req, good for tracing
	router.Use(middleware.RequestID)
	// add real client ip to request
	router.Use(middleware.RealIP)
	// add request handling logger
	// router.Use(middleware.Logger) // default solution
	//router.Use(middlewareLogger.New(logger)) // custom solution
	// if panic don't die, just 500
	router.Use(middleware.Recoverer)
	// functional urls for handlers ('/articles/{id}')
	router.Use(middleware.URLFormat)
}

func setupRoutes(router *chi.Mux) {
	router.Get("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte(time.Now().String()))
	})
}
