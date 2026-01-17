package server

import (
	"encoding/json"
	"net/http"
)

type params struct {
	Name string `json:"name"`
}

func (s *server) JoinHandler(res http.ResponseWriter, req *http.Request) {
	var params params
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&params)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	_, err = s.playerService.Add(params.Name)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
