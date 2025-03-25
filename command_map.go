package main

import (
	"fmt"
)

func commandMapf(config *Config, args ...string) error {
	locationAreas, err := config.pokeapiClient.ListLocations(config.nextURL)

	if err != nil {
		return err
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	config.nextURL = locationAreas.Next
	config.previousURL = locationAreas.Previous

	return nil
}

func commandMapb(config *Config, args ...string) error {
	if config.previousURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationAreas, err := config.pokeapiClient.ListLocations(config.previousURL)
	if err != nil {
		return err
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	config.nextURL = locationAreas.Next
	config.previousURL = locationAreas.Previous

	return nil

}
