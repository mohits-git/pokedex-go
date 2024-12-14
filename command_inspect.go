package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("missing pokemon name; usage: inspect <pokemon name>")
	}
	pokemonName := cfg.args[0]

	pokedex := cfg.pokedex

	pokemon, ok := pokedex.pokemons[pokemonName]

	if !ok {
		return errors.New(fmt.Sprintf("You don't have %s in your Pokeball", pokemonName))
	}

  fmt.Printf("Name: %s\n", pokemon.Name)
  fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
  fmt.Printf("Height: %d\n", pokemon.Height)
  fmt.Printf("Weight: %d\n", pokemon.Weight)
  fmt.Printf("Stats:\n")
  for _, stat := range pokemon.Stats {
    fmt.Printf("  - %s: %d\n", stat.Name, stat.BaseStat)
  }
  fmt.Printf("Types:\n")
  for _, t := range pokemon.Types {
    fmt.Printf("  - %s\n", t)
  }
	return nil
}
