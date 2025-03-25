package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	pokemonName := args[0]
	pokemonDetail, err := config.pokeapiClient.ListPokemonDetails(pokemonName)

	if err != nil {
		return err
	}

	if pokemonDetail.Name == "" {
		return fmt.Errorf("Pokemon does not exist")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catchPokemon(pokemonDetail.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command")
		config.pokeDex[pokemonName] = pokemonDetail
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// TODO come up with a better Algorithm for catching
func catchPokemon(baseXP int) bool {
	if rand.Intn(1000) > baseXP {
		return true
	}
	return false
}
