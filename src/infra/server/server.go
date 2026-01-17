// Package server is the package for creating the http server and it's routes
package server

import (
	"net/http"

	"github.com/jaypeenascimento/myyouquiz/src/domain/service"
)

type server struct {
	playerService *service.PlayerService
}

func NewServer(playerService *service.PlayerService) http.Handler {
	s := server{
		playerService: playerService,
	}

	router := http.NewServeMux()

	// Registering routes
	router.HandleFunc("GET /health", s.HealthHandler)
	router.HandleFunc("POST /join", s.JoinHandler)

	return router
}
