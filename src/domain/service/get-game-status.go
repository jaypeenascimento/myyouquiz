package service

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

type GameStatus struct {
	CurrentKing string          `json:"currentKing"`
	ScoreToWin  int             `json:"scoreToWin"`
	LetterQueue []entity.Letter `json:"letterQueue"`
	Players     []entity.Player `json:"players"`
}

func (s *GameService) GetGameStatus() (GameStatus, error) {
	players, err := s.playersRepo.FetchAllPlayers()
	if err != nil {
		return GameStatus{}, err
	}

	return GameStatus{
		Players: players,
	}, nil
}
