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
    return parsePokemon(cachedResp)
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

	c.cache.Add(url, data)

  return parsePokemon(data)
}

func parsePokemon(data []byte) (Pokemon, error) {
	pokemonsResp := PokemonResponse{}
	if err := json.Unmarshal(data, &pokemonsResp); err != nil {
		return Pokemon{}, err
	}

	types := []string{}
	for _, t := range pokemonsResp.Types {
		types = append(types, t.Type.Name)
	}
	stats := []Stat{}
	for _, s := range pokemonsResp.Stats {
		stats = append(
			stats,
			Stat{
				Name:     s.Stat.Name,
				BaseStat: s.BaseStat,
			},
		)
	}

	pokemon := Pokemon{
		ID:             pokemonsResp.ID,
		Name:           pokemonsResp.Name,
		BaseExperience: pokemonsResp.BaseExperience,
		Height:         pokemonsResp.Height,
		Weight:         pokemonsResp.Weight,
		Types:          types,
    Stats:          stats,
	}

	return pokemon, nil
}
