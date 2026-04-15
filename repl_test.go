package main

import(
	"fmt"
	"testing"
	"time"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 		string
		expected  []string
	}{
			{
				input:	"  hello world ",
				expected: []string{"hello", "world"},
			},
			{
				input: " Yo",
				expected: []string{"yo"},
			},
			{
				input: " helLlLlLooooo wo rl D    ",
				expected: []string{"hellllllooooo", "wo", "rl", "d"},
			},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}

func TestCache() {
	times = []time.Time

	pokeClient := pokeapi.NewClient(time.Second)
	duration := 5 * time.Second
	cache := pokecache.NewCache(duration) 

	c := &config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0",
		previous: "",
		client: pokeClient,
		cache: cache
	}
	
	fmt.Printf("%v\n", time.Now())
	commandMapf(c)
	fmt.Printf("%v\n", time.Now())
	commandMapb(c)
	fmt.Printf("%v\n", time.Now())
}
