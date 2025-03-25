package main

import (
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.pokeDex) == 0 {
		return fmt.Errorf("You have not caught any pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokeDex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
