package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name

	cachedResp, cached := c.cache.Get(url)
	if cached {
		pokemonsResp := PokemonResponse{}
		if err := json.Unmarshal(cachedResp, &pokemonsResp); err != nil {
			return Pokemon{}, err
		}

		pokemon := Pokemon{
			ID:                     pokemonsResp.ID,
			Name:                   pokemonsResp.Name,
			BaseExperience:         pokemonsResp.BaseExperience,
			LocationAreaEncounters: pokemonsResp.LocationAreaEncounters,
		}

		return pokemon, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonsResp := PokemonResponse{}
	if err = json.Unmarshal(data, &pokemonsResp); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	pokemon := Pokemon{
		ID:                     pokemonsResp.ID,
		Name:                   pokemonsResp.Name,
		BaseExperience:         pokemonsResp.BaseExperience,
		LocationAreaEncounters: pokemonsResp.LocationAreaEncounters,
	}

	return pokemon, nil
}
