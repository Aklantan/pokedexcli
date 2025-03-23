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

var commands map[string]cliCommand

func startRepl() {
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
					err := command.callback()
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
