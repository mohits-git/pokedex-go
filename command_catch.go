package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
  if len(cfg.args) < 1 {
    return errors.New("Please provide a pokemon name to catch")
  }

  // Get user pokedex and check if the pokemon is already caught
  pokedex := cfg.pokedex

  _, ok := pokedex.pokemons[cfg.args[0]]
  if ok {
    return errors.New("You already have this pokemon")
  }

  fmt.Printf("Throwing a Pokeball at %s...\n", cfg.args[0])
  pokemon, err := cfg.pokeapiClient.GetPokemon(cfg.args[0])
  if err != nil {
    return err
  }

  // Randomly decide if the pokemon is caught
  prob := rand.Intn(pokemon.BaseExperience)
  isCaught := false
  if prob >= (pokemon.BaseExperience*2/3) {
    isCaught = true
  }

  // Get the user's number of attempts for this pokemon
  attempts, ok := pokedex.userAttemps[pokemon.Name]
  if !ok {
    attempts = 1
  }
  attempts++

  if isCaught || attempts >= 5 {
    pokedex.pokemons[pokemon.Name] = pokemon
    pokedex.userAttemps[pokemon.Name] = attempts
    fmt.Printf("%s was caught!\n", pokemon.Name)
    return nil
  }

  pokedex.userAttemps[pokemon.Name] = attempts
  fmt.Printf("%s escaped!\n", pokemon.Name)

  return nil
}
