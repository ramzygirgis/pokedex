package main

import(
	"time"
	// "fmt"
	"github.com/ramzygirgis/pokedex/internal/pokeapi"
	"github.com/ramzygirgis/pokedex/internal/pokecache"
)


func main() {
	pokeClient := pokeapi.NewClient(time.Second)
	duration := 5 * time.Second
	cache := pokecache.NewCache(duration) 

	c := &config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
		client: pokeClient,
		cache: *cache,
	}

	startRepl(c)
}



/*

func main() {
	pokeClient := pokeapi.NewClient(time.Second)
	duration := 5 * time.Second
	cache := *pokecache.NewCache(duration) 

	c := &config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
		client: pokeClient,
		cache: cache,
	}
	
	fmt.Printf("%v\n", time.Now())
	commandMapf(c)
	fmt.Printf("%v\n", time.Now())
	commandMapb(c)
	fmt.Printf("%v\n", time.Now())
}

*/
