package main

import "github/Aklantan/pokedexcli/internal/pokecache"

func main() {
	config := NewConfig()
	config.Pokedex = make(map[string]PokemonProfile)
	cache := pokecache.NewCache(7)
	startRepl(config, cache)

}
