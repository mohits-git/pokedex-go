package main

import (
	"time"

	"github.com/mohits-git/pokedex-go/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
    args:          []string{},
	}
	startRepl(cfg)
}
