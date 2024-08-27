// Created by Anh Cao on 27.08.2024.

package main

import (
	"log"
	"net/http"

	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/api/handlers"
	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/api/middleware"
	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/api/routes"
	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/cache"
	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/logger"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize cache
	cache.NewCache()

	// Initial new router
	r := mux.NewRouter()
	for _, endpoint := range routes.Endpoints {
		r.HandleFunc(endpoint.Path, endpoint.Handler).Methods(endpoint.Method)
	}
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.NotAllowed)
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	// Middleware
	r.Use(middleware.Logger)

	// Start server
	logger.Logger.Info("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
