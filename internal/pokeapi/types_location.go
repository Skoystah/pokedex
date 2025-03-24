package pokeapi

// Result JSON of location-area API call
// Created by using https://mholt.github.io/json-to-go/
type LocationAreaResult struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
