// Package server is the package for creating the http server and it's routes
package server

import "net/http"

func NewServer() http.Handler {
	router := http.NewServeMux()

	// Registering routes
	router.HandleFunc("GET /health", HealthHandler)

	return router
}
