package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	ALLOWED_ORIGIN := os.Getenv("ALLOWED_ORIGIN")
	log.Printf("ALLOWED_ORIGIN: %s\n", ALLOWED_ORIGIN)
	// Specfi which origins can have cors acces
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173", "http://localhost", "http://127.0.0.1", ALLOWED_ORIGIN},
		AllowedMethods:   []string{http.MethodDelete, http.MethodPut, http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)
	mux.Post("/", app.Broker)
	mux.Post("/handle", app.HandleSubmission)
	mux.Post("/log-grpc", app.LogViaGRPC)
	return mux
}
