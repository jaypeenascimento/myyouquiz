package server

import (
	"net/http"
)

func (s *server) StartHandler(res http.ResponseWriter, req *http.Request) {
	err := s.gameService.StartGame()
	if err != nil {
		sendInternalServerError(res, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
