package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func routes() http.Hander {
    mux := chi.NewRouter

    //specify who is allowed to connect
    mux.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"https://*", "http://*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredenctials: true,
        MaxAge: 300,
    }))

    mux>use(middleware.Heatbeat("/ping"))
}