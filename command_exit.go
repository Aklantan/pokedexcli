package main

import (
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
	"os"
)

func commandExit(config *Config, cache *pokecache.Cache, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	err := fmt.Errorf("program exited")
	os.Exit(0)
	return err
}
