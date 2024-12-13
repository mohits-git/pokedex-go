package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func processLocationAreasCommand(url string, cfg *config) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	result := LocationAreasResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	cfg.nextLocationUrl = result.Next
	cfg.prevLocationUrl = result.Previous
	for _, locationArea := range result.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapf(cfg *config) error {
	locationAreaApiUrl := "https://pokeapi.co/api/v2/location-area"
	url := cfg.nextLocationUrl
	if url == nil {
    url = &locationAreaApiUrl
	}
	return processLocationAreasCommand(*url, cfg)
}

func commandMapb(cfg *config) error {
	url := cfg.prevLocationUrl
	if url == nil {
		return errors.New("you're on the first page")
	}
	return processLocationAreasCommand(*url, cfg)
}
