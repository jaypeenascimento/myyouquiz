package service

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/jaypeenascimento/myyouquiz/src/domain/entity"
)

const DefaultScoreToWin int = 10

var ErrCannotStartGame = errors.New("cannot start game due to unsufficient players in room")

func (s GameService) StartGame() error {
	players, err := s.playersRepo.FetchAllPlayers()
	if err != nil || len(players) <= 1 {
		return ErrCannotStartGame
	}

	randomInteger := rand.IntN(len(players))
	chosen := players[randomInteger]

	var game entity.Game
	game.KingPlayer = chosen.Name
	game.ScoreToWin = DefaultScoreToWin

	_, err = s.gameRepository.SetGame(game)
	if err != nil {
		return fmt.Errorf("failed to update game data at game repository: %w", err)
	}

	return nil
}
