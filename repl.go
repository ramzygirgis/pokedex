package main

import "strings"

func cleanInput(s string) []string {
	/*
	given a string s, cleanInput produces a slice of strings containing the "words" of s
	according to whitespace. 
	All words are converted to lowercase, and all leading or trailing whitespace is trimmed.
	*/
	cleanStrings := make([]string, 0)
	l := 0
	for r := 0; r <len(s); r++ {
		if string(s[r]) == " " {
			if l < r {
				cleanStrings = append(cleanStrings, strings.ToLower(s[l:r]))
			}	
			l = r + 1
			continue
		}
	}
	if l < len(s) {
		cleanStrings = append(cleanStrings, strings.ToLower(s[l:len(s)]))
	}
	return cleanStrings
}
