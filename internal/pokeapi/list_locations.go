package pokeapi

import (
	"encoding/json"
	//"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResult, error) {
	fullURL := baseURL + "location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var data []byte
	if cachedEntry, exists := c.cache.Get(fullURL); exists {
		//fmt.Printf("retrieving from cache %v\n", fullURL)
		data = cachedEntry
	} else {
		res, err := c.httpClient.Get(fullURL)
		if err != nil {
			return LocationAreaResult{}, err
		}
		defer res.Body.Close()

		//Decode JSON results
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResult{}, err
		}

		c.cache.Add(fullURL, data)
	}

	var locationAreas LocationAreaResult

	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreaResult{}, err
	}

	return locationAreas, nil
}
