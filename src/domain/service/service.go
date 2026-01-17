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
	playersRepo repository.PlayerRepository
}

func NewGameService(playerRepo repository.PlayerRepository) *GameService {
	return &GameService{
		playersRepo: playerRepo,
	}
}
