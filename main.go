package main

import (
	"time"

	"github.com/mohits-git/pokedex-go/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(10*time.Second, 5*time.Minute),
		args:          []string{},
		pokedex: &userPokedex{
			pokemons: map[string]pokeapi.Pokemon{},
      userAttemps: map[string]int{},
		},
	}
	startRepl(cfg)
}
