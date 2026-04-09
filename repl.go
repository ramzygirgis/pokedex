package main

import(
	"strings"
	"github.com/ramzygirgis/pokedex/internal/pokeapi"
	"fmt"
	"os"
	"bufio"
)

type config struct {
	next string
	previous string
	client pokeapi.Client
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			input := scanner.Text()
			tokens := cleanInput(input)
			commandName := ""
			if len(tokens) != 0 {
				commandName = tokens[0]
			}
			
			if cmd, ok := getCommands()[commandName]; ok {
				err := cmd.callback(c)
				if err != nil {
		  		fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
