package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {

	pokeClient := pokeapi.NewClient(time.Second * 5)

	config := &Config{pokeapiClient: pokeClient}

	startRepl(config)
}
