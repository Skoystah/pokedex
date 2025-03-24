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
		handler, exists := cliCommands[iptCommand]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := handler.callback(config)
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
	callback    func(*Config) error
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
	}
}
