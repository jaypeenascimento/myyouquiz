package service

func (s *GameService) GetRoleOfPlayer(name string) (string, error) {
	game, err := s.gameRepository.GetGame()
	if err != nil {
		return "", err
	}

	var role string
	if name == game.KingPlayer {
		role = "King"
	} else {
		role = "Peasant"
	}

	return role, nil
}
