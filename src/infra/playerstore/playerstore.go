// Package playerstore groups all functionality to manage Players at persistance layer.
package playerstore

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

type store struct {
	players map[string]entity.Player
}

func NewPlayerstore() *store {
	return &store{
		players: map[string]entity.Player{},
	}
}
