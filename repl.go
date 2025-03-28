package main

import (
	"bufio"
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, cache *pokecache.Cache, parameter string) error
}

var commands map[string]cliCommand

func startRepl(config *Config, cache *pokecache.Cache) {
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
		"explore": {
			name:        "explore",
			description: "Displays the pokemon that can be found in the area. Add location as a parameter to this command",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon. Add pokemon as a parameter to this command",
			callback:    commandCatch,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	parameter := ""
	for {
		fmt.Print("Pokedex> ")
		if scanner.Scan() {
			line := scanner.Text()
			input := cleanInput(line)
			if len(input) > 0 {
				command, exists := commands[input[0]]
				if len(input) > 1 {
					parameter = input[1]
				}
				if exists {
					err := command.callback(config, cache, parameter)
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
