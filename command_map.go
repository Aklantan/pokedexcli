package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextLocationAreaURL != nil {
		url = *config.NextLocationAreaURL
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	areas := LocationList{}
	config.NextLocationAreaURL = &areas.NextLocation
	config.PreviousLocationAreaURL = &areas.PreviousLocation
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range areas.Locations {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(config *Config) error {
	if *config.PreviousLocationAreaURL == "" {
		return fmt.Errorf("you're on the first page")
	}
	url := *config.PreviousLocationAreaURL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	areas := LocationList{}
	config.NextLocationAreaURL = &areas.NextLocation
	config.PreviousLocationAreaURL = &areas.PreviousLocation
	err = json.Unmarshal(body, &areas)
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range areas.Locations {
		fmt.Println(location.Name)
	}

	return nil
}
