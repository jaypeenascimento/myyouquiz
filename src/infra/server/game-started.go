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
	var params gameStartedParams

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var response gameStartedResponse
	encoder := json.NewEncoder(res)

	if s.gameService.IsGameStarted() {

		role, err := s.gameService.GetRoleOfPlayer(params.Name)
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
