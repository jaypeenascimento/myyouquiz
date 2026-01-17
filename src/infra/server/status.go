package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) StatusHandler(res http.ResponseWriter, req *http.Request) {
	status, err := s.gameService.GetGameStatus()
	if err != nil {
		sendInternalServerError(res, err)
		return
	}

	encoder := json.NewEncoder(res)
	if err = encoder.Encode(status); err != nil {
		sendInternalServerError(res, err)
		return
	}
}

func sendInternalServerError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintf(res, "Unexpected error: %s", err.Error())
}
