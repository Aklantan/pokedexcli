package main

import (
	"encoding/json"
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type EncounterList struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

func commandExplore(config *Config, cache *pokecache.Cache, parameter string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + parameter
	fmt.Printf("Exploring %s\n", parameter)
	exploreHelper(url, config, cache)
	return nil

}

func exploreHelper(url string, config *Config, cache *pokecache.Cache) error {
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
	var pokemons EncounterList
	{
	}

	err := json.Unmarshal(body, &pokemons)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons.Encounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
