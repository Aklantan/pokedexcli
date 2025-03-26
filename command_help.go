package main

import (
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil

}
