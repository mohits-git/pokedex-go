package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.args) < 1 {
		return errors.New("please provide a location name")
	}
	locationName := cfg.args[0]

	pokemons, err := cfg.pokeapiClient.GetLocationPokemons(locationName)
	if err != nil {
		return err
	}

  fmt.Printf("Exploring %s...\n", locationName)
  fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
