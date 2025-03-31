package main

import (
	"encoding/json"
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
	"io"
	"math/rand"
	"net/http"
)

type PokemonProfile struct {
	Name           string     `json:"name"`
	BaseExperience int        `json:"base_experience"`
	ID             int        `json:"id"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Types          []PokeType `json:"types"`
	Stats          []PokeStat `json:"stats"`
}

type PokeType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type PokeStat struct {
	Stat struct {
		Name string `json:"name"`
	} `json:"stat"`
	Value int `json:"base_stat"`
}

func commandCatch(config *Config, cache *pokecache.Cache, parameter string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + parameter
	fmt.Printf("Exploring %s\n", parameter)
	catchHelper(url, config, cache)
	return nil

}

func catchHelper(url string, config *Config, cache *pokecache.Cache) error {
	var body []byte
	cached_data, exists := cache.Get(url)
	if exists {
		body = cached_data
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		cache.Add(url, body)

	}
	var pokemon PokemonProfile

	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := 255 - pokemon.BaseExperience
	pokeball := rand.Intn(255)
	if pokeball < chance {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		config.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(config *Config, cache *pokecache.Cache, parameter string) error {
	_, exists := config.Pokedex[parameter]
	if exists {
		fmt.Printf("Name: %s\n", config.Pokedex[parameter].Name)
		fmt.Printf("ID: %d\n", config.Pokedex[parameter].ID)
		fmt.Printf("Height: %d\n", config.Pokedex[parameter].Height)
		fmt.Printf("Weight: %d\n", config.Pokedex[parameter].Weight)
		fmt.Println("Types:")
		for _, poketype := range config.Pokedex[parameter].Types {
			fmt.Printf("    - %s\n", poketype.Type.Name)
		}
		fmt.Println("Stats:")
		for _, pokestat := range config.Pokedex[parameter].Stats {
			fmt.Printf("    - %s : % d\n", pokestat.Stat.Name, pokestat.Value)
		}

	} else {
		fmt.Printf("%s is not in your Pokedex yet.\n", parameter)
	}

	return nil
}

func commandPokedex(config *Config, cache *pokecache.Cache, parameter string) error {
	fmt.Println("Your Pokedex - ")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("    - %s\n", pokemon.Name)
	}
	fmt.Println("You can use the inspect command to look at details of these pokemon.")
	return nil

}
