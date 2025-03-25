package main

import (
	"fmt"
)

func commandInspect(config *Config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	pokemonName := args[0]

	pokemon, exists := config.pokeDex[pokemonName]

	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %s\n", typ.Type.Name)
	}

	return nil
}
