// Package service contains the domain services provided
package service

import "github.com/jaypeenascimento/myyouquiz/src/domain/repository"

type PlayerService struct {
	playersRepo repository.PlayerRepository
}

func NewService(playerRepo repository.PlayerRepository) *PlayerService {
	return &PlayerService{
		playersRepo: playerRepo,
	}
}

type GameService struct {
	playersRepo    repository.PlayerRepository
	gameRepository repository.GameRepository
}

func NewGameService(playerRepo repository.PlayerRepository, gameRepository repository.GameRepository) *GameService {
	return &GameService{
		playersRepo:    playerRepo,
		gameRepository: gameRepository,
	}
}
