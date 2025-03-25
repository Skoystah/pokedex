package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListPokemonDetails(pokemonName string) (PokemonDetailResult, error) {
	fullURL := baseURL + "pokemon/" + pokemonName

	var data []byte
	if cachedEntry, exists := c.cache.Get(fullURL); exists {
		//fmt.Printf("retrieving from cache %v\n", fullURL)
		data = cachedEntry
	} else {
		res, err := c.httpClient.Get(fullURL)
		if err != nil {
			return PokemonDetailResult{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return PokemonDetailResult{}, fmt.Errorf("Pokemon does not exist")
		}

		// Decode JSON results
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonDetailResult{}, err
		}
		c.cache.Add(fullURL, data)
	}

	var pokemonDetail PokemonDetailResult

	err := json.Unmarshal(data, &pokemonDetail)
	if err != nil {
		return PokemonDetailResult{}, err
	}

	return pokemonDetail, nil
}
