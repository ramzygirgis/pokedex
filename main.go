package main

import "fmt"
import "bufio"
import "os"


func main() {
	c := config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
	}

	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: 		"map",
			description: "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locations",
			callback: commandMapb,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for {
		if scanner.Scan() {
			input := scanner.Text() // might need to "forward declare" token
			tokens := cleanInput(input)
			if len(tokens) != 0 {
				if cmd, ok := commandRegistry[tokens[0]]; ok {
				err := cmd.callback(&c)
				if err != nil {
		  		fmt.Println(err)
				}
				} else {
					fmt.Println("Unknown command")
				}
			} else {
				fmt.Println("Unknown command")
			}
			fmt.Print(prompt)
		}
	}
}

