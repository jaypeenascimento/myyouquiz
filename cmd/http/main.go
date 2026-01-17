package main

import (
	"fmt"
	"net/http"

	"github.com/jaypeenascimento/myyouquiz/src/domain/service"
	"github.com/jaypeenascimento/myyouquiz/src/infra/gamestore"
	"github.com/jaypeenascimento/myyouquiz/src/infra/playerstore"
	"github.com/jaypeenascimento/myyouquiz/src/infra/server"
)

func main() {
	fmt.Println("Server started!")

	playerRepo := playerstore.NewPlayerstore()
	gameRepo := gamestore.NewGameStore()

	playerService := service.NewService(playerRepo)
	gameService := service.NewGameService(playerRepo, gameRepo)

	mux := server.NewServer(playerService, gameService)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(fmt.Sprintf("Server crashed with error: %s", err.Error()))
	}
}
