package main

import "time"
import "github.com/ramzygirgis/pokedex/internal/pokeapi"


func main() {
	pokeClient := pokeapi.NewClient(time.Second)

	c := &config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
		client: pokeClient,
	}

	startRepl(c)
}

