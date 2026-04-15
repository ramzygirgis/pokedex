package main

import(
	"time"
	"github.com/ramzygirgis/pokedex/internal/pokeapi"
	"github.com/ramzygirgis/pokedex/internal/pokecache"
)

/*
func main() {
	pokeClient := pokeapi.NewClient(time.Second)
	duration := 5 * time.Second
	cache := pokecache.NewCache(duration) 

	c := &config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
		client: pokeClient,
		cache: cache
	}

	startRepl(c)
}
*/

func main() {
	TestCache()
}
