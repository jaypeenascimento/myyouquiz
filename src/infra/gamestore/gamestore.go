// Package gamestore groups all functions related to game data persistance
package gamestore

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

type store struct {
	data *entity.Game
}

func NewGameStore() *store {
	return &store{}
}
