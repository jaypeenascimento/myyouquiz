// Package server is the package for creating the http server and it's routes
package server

import (
	"net/http"

	"github.com/jaypeenascimento/myyouquiz/src/domain/service"
)

type server struct {
	playerService *service.PlayerService
	gameService   *service.GameService
}

func NewServer(playerService *service.PlayerService, gameService *service.GameService) http.Handler {
	s := server{
		playerService: playerService,
		gameService:   gameService,
	}

	router := http.NewServeMux()

	// Registering routes
	router.HandleFunc("GET /health", s.HealthHandler)
	router.HandleFunc("POST /join", s.JoinHandler)
	router.HandleFunc("GET /game-status", s.StatusHandler)
	router.HandleFunc("GET /game-started/{name}", s.GameStartedHandler)
	router.HandleFunc("POST /start", s.StartHandler)

	return router
}
