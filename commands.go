package main

import "fmt"
import "os"
import "net/http"
import "encoding/json"

type cliCommand struct {
	name	string
	description string
	callback func(*config) error
}

type config struct {
	next string
	current string
	previous string
}

type locationArea struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}


var commandRegistry map[string]cliCommand

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range commandRegistry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(c *config) error {
	req, err := http.NewRequest("GET", c.next, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	
	var l locationArea
	if err := json.NewDecoder(resp.Body).Decode(&l); err != nil {
		return err
	}

	for i := 0; i < len(l.Results); i++ {
		fmt.Println(l.Results[i].Name)
	}

	c.previous = c.next
	c.next = l.Next
	return nil

}

func commandMapb(c *config) error {
	if c.previous == "" {
		fmt.Println("No previous locations\n")
		return nil
	}
	req, err := http.NewRequest("GET", c.previous, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	
	var l locationArea
	if err := json.NewDecoder(resp.Body).Decode(&l); err != nil {
		return err
	}

	for i := 0; i < len(l.Results); i++ {
		fmt.Println(l.Results[i].Name)
	}
	
	c.next = c.previous
	c.previous = l.Previous
	return nil

}
