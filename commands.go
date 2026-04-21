package main

import "fmt"
import "os"
import "github.com/ramzygirgis/pokedex/internal/pokeapi"


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

	l, ok := c.cache.Get(url)
	if !ok {
		var err error
		l, err = c.client.PokeapiCall(url)
		if err != nil {
			return err
		}
		c.cache.Add(url, l)
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
	
	l, ok := c.cache.Get(url)
	if !ok {
		var err error
		l, err = c.client.PokeapiCall(url)
		if err != nil {
			return err
		}
		c.cache.Add(url, l)
	}

	navUrl, err := pokeapi.ListLocations(false, l)
	if err != nil {
		return err
	}
	
	c.next = url
	c.previous = navUrl
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
	}
}
