package main

import (
	"errors"
	"fmt"
)

func processLocationAreasCommand(pageUrl *string, cfg *config) error {
	locationsList, err := cfg.pokeapiClient.ListLocations(pageUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationsList.Next
	cfg.prevLocationUrl = locationsList.Previous
	for _, locationArea := range locationsList.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapf(cfg *config) error {
	return processLocationAreasCommand(cfg.nextLocationUrl, cfg)
}

func commandMapb(cfg *config) error {
	url := cfg.prevLocationUrl
	if url == nil {
		return errors.New("you're on the first page")
	}
	return processLocationAreasCommand(url, cfg)
}
