package main

import "fmt"

func commandPokedex(cfg *config) error {
  pokedex := cfg.pokedex

  fmt.Println("Your Pokedex:")
  for p := range pokedex.pokemons {
    fmt.Printf(" - %s\n", p)
  }

  return nil
}
