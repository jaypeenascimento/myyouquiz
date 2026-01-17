package service

func (s *GameService) IsGameStarted() bool {
	_, err := s.gameRepository.GetGame()
	return err == nil

	// If started, no error. If not started, haves error.
}
