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
	callback    func(config *Config) error
}

var commands map[string]cliCommand

func startRepl(config *Config) {
	commands = map[string]cliCommand{
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
			description: "Displays the areas available in groups of 20, subsequent calls move to the next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 areas available",
			callback:    commandMapB,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		if scanner.Scan() {
			line := scanner.Text()
			input := cleanInput(line)
			if len(input) > 0 {
				command, exists := commands[input[0]]
				if exists {
					err := command.callback(config)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}

			}
		}
	}

}

func cleanInput(text string) []string {
	lwrString := strings.ToLower(text)
	words := strings.Fields(strings.ToLower(lwrString))

	return words

}
