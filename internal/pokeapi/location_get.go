package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemons(locationName string) ([]string, error) {
	url := baseUrl + "/location-area/" + locationName

	cachedResp, cached := c.cache.Get(url)
	if cached {
		pokemonsResp := GetLocationResponse{}
		if err := json.Unmarshal(cachedResp, &pokemonsResp); err != nil {
			return []string{}, err
		}

		pokemons := []string{}
		for _, pokemon := range pokemonsResp.PokemonEncounters {
			pokemons = append(pokemons, pokemon.Pokemon.Name)
		}

		return pokemons, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	c.cache.Add(url, data)

	pokemonsResp := GetLocationResponse{}
	err = json.Unmarshal(data, &pokemonsResp)
	if err != nil {
		return []string{}, err
	}

	pokemons := []string{}
	for _, pokemon := range pokemonsResp.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}

	return pokemons, nil
}
