package main

import "fmt"
import "bufio"
import "os"


func main() {
	// declare "globals". would be sick if i can find a way to put this in another file
	//
	prompt := "Pokedex > "
	commandRegistry := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
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
				err := cmd.callback()
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

