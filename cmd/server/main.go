package main

import (
	"LearnTest/internal/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.HandleFunc("/cafe", handlers.MainHandle)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panic("start server error")
	}

}
