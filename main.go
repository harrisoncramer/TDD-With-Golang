package main

import (
	"fmt"
	"log"
	"net/http"
	"tdd_with_golang/server"
)

type InMemoryScoreDB struct{}

const port = 3000

func (i *InMemoryScoreDB) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &server.PlayerServer{
		Store: &InMemoryScoreDB{},
	}

	handler := http.HandlerFunc(server.ServeHTTP)
	fmt.Printf("Listening on port :%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		log.Fatal(err)
	}

}
