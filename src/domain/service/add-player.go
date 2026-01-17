package service

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

func (s *PlayerService) Add(name string) (entity.Player, error) {
	player, _ := s.playersRepo.FindPlayer(name)

	if player.Name == name {
		return player, nil
	}

	newPlayer := *entity.NewPlayer(name)
	_, err := s.playersRepo.PutPlayer(newPlayer)
	if err != nil {
		return entity.Player{}, err
	}

	return newPlayer, nil
}
