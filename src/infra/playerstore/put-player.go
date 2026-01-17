package playerstore

import (
	"errors"

	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

var ErrInvalidPlayerName = errors.New("cannot put a player with empty name")

func (s *store) PutPlayer(player entity.Player) (entity.Player, error) {
	if player.Name == "" {
		return entity.Player{}, ErrInvalidPlayerName
	}

	s.players[player.Name] = player

	return s.players[player.Name], nil
}
