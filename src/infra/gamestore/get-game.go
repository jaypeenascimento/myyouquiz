package gamestore

import (
	"errors"

	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

var ErrNotStartedGame = errors.New("game was not started, cannot get data of it")

func (s *store) GetGame() (entity.Game, error) {
	if s.data == nil {
		return entity.Game{}, ErrNotStartedGame
	}

	return *s.data, nil
}
