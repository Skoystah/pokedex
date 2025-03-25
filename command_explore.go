package main

import (
	"fmt"
)

func commandExplore(config *Config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("no arguments provided")
	}

	locationArea := args[0]
	locationAreaDetail, err := config.pokeapiClient.ListLocationDetails(locationArea)

	if err != nil {
		return err
	}

	if len(locationAreaDetail.PokemonEncounters) == 0 {
		return fmt.Errorf("No Pokemon at this location! :(")
	}

	for _, result := range locationAreaDetail.PokemonEncounters {
		fmt.Println(result.Pokemon.Name)
	}

	return nil
}
