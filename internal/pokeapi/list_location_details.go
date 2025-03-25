package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocationDetails(locationArea string) (LocationAreaDetailResult, error) {
	fullURL := baseURL + "location-area/" + locationArea

	var data []byte
	if cachedEntry, exists := c.cache.Get(fullURL); exists {
		//fmt.Printf("retrieving from cache %v\n", fullURL)
		data = cachedEntry
	} else {
		res, err := c.httpClient.Get(fullURL)
		if err != nil {
			return LocationAreaDetailResult{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return LocationAreaDetailResult{}, fmt.Errorf("Nothing found")
		}

		// Decode JSON results
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetailResult{}, err
		}
		c.cache.Add(fullURL, data)
	}

	var locationAreaDetail LocationAreaDetailResult

	err := json.Unmarshal(data, &locationAreaDetail)
	if err != nil {
		return LocationAreaDetailResult{}, err
	}

	return locationAreaDetail, nil
}
