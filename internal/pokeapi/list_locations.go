package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResult, error) {
	fullURL := baseURL + "location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreaResult{}, fmt.Errorf("error Get location res", err)
	}
	defer res.Body.Close()

	//Decode JSON results
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResult{}, fmt.Errorf("error reading data", err)
	}

	//decoder := json.NewDecoder(res.Body)
	var locationAreas LocationAreaResult

	//err = decoder.Decode(&locationAreas)
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationAreaResult{}, fmt.Errorf("error unmarshaling res", err)
	}

	return locationAreas, nil
}
