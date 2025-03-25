package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {

	pokeClient := pokeapi.NewClient(time.Second*5, time.Second*60)

	config := &Config{pokeapiClient: pokeClient,
		pokeDex: map[string]pokeapi.PokemonDetailResult{},
	}

	startRepl(config)
}
