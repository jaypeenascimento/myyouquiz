package playerstore

import (
	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

func (s *store) FetchAllPlayers() ([]entity.Player, error) {
	var players []entity.Player
	for _, value := range s.players {
		players = append(players, value)
	}

	return players, nil
}
