package main

import "fmt"
import "os"

type cliCommand struct {
	name	string
	description string
	callback func() error
}



func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
