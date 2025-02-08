package main

import (
	"go-first-big-project/api"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Starting server...", "version", "1.0.0")
	if err := run(); err != nil {
		slog.Error("failed to run", "error", err)
		os.Exit(1)
	}
	slog.Info("All systems offline")
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env file", "error", err)
		os.Exit(1)
	}
	apiKey := os.Getenv("OMDB_KEY")

	handler := api.NewHandler(apiKey)
	s := http.Server{
		Addr:                         ":8080",
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  time.Minute,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
