package playerstore

import (
	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

func (s *store) FindPlayer(name string) (entity.Player, error) {
	return s.players[name], nil
}
