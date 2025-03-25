package main

import (
	"encoding/json"
	"fmt"
	"github/Aklantan/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationList struct {
	Locations        []Location `json:"results"`
	NextLocation     string     `json:"next"`
	PreviousLocation string     `json:"previous"`
}

func commandMap(config *Config, cache *pokecache.Cache) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextLocationAreaURL != nil {
		url = *config.NextLocationAreaURL
	}
	mapHelper(url, config, cache)

	return nil
}

func commandMapB(config *Config, cache *pokecache.Cache) error {
	if config.PreviousLocationAreaURL == nil || *config.PreviousLocationAreaURL == "" {
		return fmt.Errorf("you're on the first page")
	}
	url := *config.PreviousLocationAreaURL
	mapHelper(url, config, cache)
	return nil
}

func mapHelper(url string, config *Config, cache *pokecache.Cache) error {
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
	areas := LocationList{}

	err := json.Unmarshal(body, &areas)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	config.NextLocationAreaURL = &areas.NextLocation
	config.PreviousLocationAreaURL = &areas.PreviousLocation
	for _, location := range areas.Locations {
		fmt.Println(location.Name)
	}

	return nil
}
