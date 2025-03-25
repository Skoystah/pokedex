package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	pokeDex       map[string]pokeapi.PokemonDetailResult
	previousURL   *string
	nextURL       *string
}

func startRepl(config *Config) {
	cliCommands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		ipt := scanner.Text()
		cleanedIpt := cleanInput(ipt)
		if len(cleanedIpt) == 0 {
			continue
		}

		iptCommand := cleanedIpt[0]

		var args []string
		if len(cleanedIpt) > 1 {
			args = cleanedIpt[1:]
		}
		handler, exists := cliCommands[iptCommand]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := handler.callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)
	return res
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"mapf": {
			name:        "mapf",
			description: "shows the next page of location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "shows the previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "shows the pokemon that can be found at given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "tries to catch a given Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
