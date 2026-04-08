package main

import "fmt"
import "bufio"
import "os"

prompt := "Pokedex"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt + ">")
		if scanner.Scan() {
			token := scanner.Text() // might need to "forward declare" token
		}


	}
}

