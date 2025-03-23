package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	previousURL any
	nextURL     string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows a list of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous list of location areas",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	cliCommands := getCommands()
	config := Config{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		ipt := scanner.Text()
		cleanedIpt := cleanInput(ipt)
		if len(cleanedIpt) == 0 {
			//            fmt.Print("\n")
			continue
		}

		iptCommand := cleanedIpt[0]
		handler, exists := cliCommands[iptCommand]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := handler.callback(&config)
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
