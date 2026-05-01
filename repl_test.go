package main

import(
	"testing"
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

