package server

import (
	"encoding/json"
	"net/http"
)

type gameStartedParams struct {
	Name string `json:"name"`
}

type gameStartedResponse struct {
	Started bool   `json:"started"`
	Role    string `json:"role,omitempty"`
}

func (s *server) GameStartedHandler(res http.ResponseWriter, req *http.Request) {
	playerName := req.PathValue("name")

	var response gameStartedResponse
	encoder := json.NewEncoder(res)

	if s.gameService.IsGameStarted() {

		role, err := s.gameService.GetRoleOfPlayer(playerName)
		if err != nil {
			sendInternalServerError(res, err)
			return
		}

		response.Started = true
		response.Role = role

	} else {
		response.Started = false
	}

	_ = encoder.Encode(response)
}
