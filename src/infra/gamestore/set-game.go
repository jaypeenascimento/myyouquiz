package gamestore

import (
	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

func (s *store) SetGame(game entity.Game) (entity.Game, error) {
	s.data = &entity.Game{}

	s.data.KingPlayer = game.KingPlayer
	s.data.ScoreToWin = game.ScoreToWin
	s.data.LetterQueue = game.LetterQueue

	return *s.data, nil
}
