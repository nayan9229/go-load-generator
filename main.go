package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/nayan9229/go-load-generator/handlers"
	"github.com/nayan9229/go-load-generator/load"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Post("/", handlers.Make(handlers.HandleHomePost))
	router.Get("/login", handlers.Make(handlers.HandleLoginIndex))
	router.Get("/result/{job_id}", handlers.Make(handlers.HandleJobDetails))
	router.Handle("/*", public())

	go load.StartJobProcessor()

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
