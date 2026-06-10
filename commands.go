package main

import (
	"fmt"
	"os"
	"math/rand"
	"github.com/ramzygirgis/pokedex/internal/pokeapi"
)


type cliCommand struct {
	name	string
	description string
	callback func(*config) error
}


func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
		return nil
}


func commandHelp(c *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}


func commandMapf(c *config) error {
	url := c.next

	l, ok := c.locationsCache.Get(url)
	if !ok {
		var err error
		l, err = c.client.LocationsCall(url)
		if err != nil {
			return err
		}
		c.locationsCache.Add(url, l)
	}

	navUrl, err := pokeapi.ListLocations(true, l)
	if err != nil {
		return err
	}
	
	c.previous = url
	c.next = navUrl
	return nil
}


func commandMapb(c *config) error {
	url := c.previous
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	
	l, ok := c.locationsCache.Get(url)
	if !ok {
		var err error
		l, err = c.client.LocationsCall(url)
		if err != nil {
			return err
		}
		c.locationsCache.Add(url, l)
	}

	navUrl, err := pokeapi.ListLocations(false, l)
	if err != nil {
		return err
	}
	
	c.next = url
	c.previous = navUrl
	return nil
}

func commandExplore(c *config) error {
	name := c.name
	if name == "" {
		fmt.Println("No pokemon name provided")
		return nil
	}

	data, ok := c.namesCache.Get(name)
	if !ok {
		var err error
		data, err = c.client.ExploreCall(name)
		if err != nil {
			return err
		}
		c.namesCache.Add(name, data)
	}

	err := pokeapi.ListEncounters(name, data)
	if err != nil {
		return err
	}
	return nil
}


func commandCatch(c *config) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", c.name)
	pokemon, err := c.client.PokemonCall(c.name)
	if err != nil {
		return err
	}

	p := pokeapi.CatchSuccessProbability(pokemon)
	num := rand.Float64()
	if p < 1 - num {
		fmt.Printf("Pokeball missed. Try again!\n")
		return nil
	}
	c.pokedex[pokemon.Name] = pokemon
	fmt.Printf("Pokemon caught! %s added to pokedex\n", pokemon.Name)
	return nil
}


func commandInspect(c *config) error {
	pokemon, exists := c.pokedex[c.name]
	if !exists {
		fmt.Printf("you have not caught that pokemon")
		return nil
	}
	pokemon.Print()
	return nil
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name: "explore",
			description: "Explore a location. Requires `locationName` argument",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a pokemon. Requires `pokemonName` argument",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a pokemon. Requires `pokemonName` argument",
			callback: commandInspect,
		},
	}
}
