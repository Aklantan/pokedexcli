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
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
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
	fmt.Printf("Throwing Pokeball at %s...\n", pokemon.Name)

	chance := 255 - pokemon.BaseExperience
	pokeball := rand.Intn(255)
	if pokeball < chance {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		config.Pokedex[pokemon.Name] = pokemon
		fmt.Print(config.Pokedex)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}
