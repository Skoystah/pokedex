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
	callback    func() error
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
	}
}

func startRepl() {
	cliCommands := getCommands()

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

		err := handler.callback()
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
