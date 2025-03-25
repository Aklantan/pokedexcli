package main

import "github/Aklantan/pokedexcli/internal/pokecache"

func main() {
	config := NewConfig()
	cache := pokecache.NewCache(7)
	startRepl(config, cache)

}
