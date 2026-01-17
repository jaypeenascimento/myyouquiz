package server

import (
	"fmt"
	"net/http"
)

func (s *server) HealthHandler(res http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprint(res, "I'm Alive")
	res.WriteHeader(http.StatusOK)

	if err != nil {
		panic("Falha ao escrever na request de health! Como assim?!")
	}
}
