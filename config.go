package main

type Config struct {
	NextLocationAreaURL     *string
	PreviousLocationAreaURL *string
	Pokedex                 map[string]PokemonProfile
}

func NewConfig() *Config {
	return &Config{}
}
