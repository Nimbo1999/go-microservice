package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewMux()
	allowedOrigins := os.Getenv("ALLOWED_ORIGIN")
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173", allowedOrigins},
		AllowedMethods:   []string{http.MethodDelete, http.MethodPut, http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Post("/authenticate", app.Authenticate)
	return mux
}
