package main

import (
	"fmt"
	"net/http"

	"github.com/jaypeenascimento/myyouquiz/src/infra/server"
)

func main() {
	fmt.Println("Server started!")

	mux := server.NewServer()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(fmt.Sprintf("Server crashed with error: %s", err.Error()))
	}
}
