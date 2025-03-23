package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *Config) error {
	var locationAreaURL string

	if config.nextURL > "" {
		locationAreaURL = config.nextURL
	} else {
		baseURL := "https://pokeapi.co/api/v2/"
		resource_locationArea := "location-area"

		locationAreaURL = baseURL + resource_locationArea
	}

	client := &http.Client{}

	//GET results
	res, err := client.Get(locationAreaURL)
	if err != nil {
		return fmt.Errorf("error Get location res", err)
	}
	defer res.Body.Close()

	//Decode JSON results
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data", err)
	}

	//decoder := json.NewDecoder(res.Body)
	var locationAreas LocationAreaResult

	//err = decoder.Decode(&locationAreas)
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshaling res", err)
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	config.nextURL = locationAreas.Next
	config.previousURL = locationAreas.Previous

	return nil
}

func commandMapb(config *Config) error {
	var locationAreaURL string

	if config.previousURL != nil {
		locationAreaURL = config.previousURL.(string)
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	client := &http.Client{}

	//GET results
	res, err := client.Get(locationAreaURL)
	if err != nil {
		return fmt.Errorf("error Get location res", err)
	}
	defer res.Body.Close()

	//Decode JSON results
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading data", err)
	}

	//decoder := json.NewDecoder(res.Body)
	var locationAreas LocationAreaResult

	//err = decoder.Decode(&locationAreas)
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return fmt.Errorf("error unmarshaling res", err)
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	config.nextURL = locationAreas.Next
	config.previousURL = locationAreas.Previous

	return nil
}
