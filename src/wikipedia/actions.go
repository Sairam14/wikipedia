package main

import (
	"fmt"
)

var ignoreTokens = []string{
	"are",
	"?",
	"is",
	"the",
	"of",
	"to",
	"some",
	"do",
	"which",
	"what",
	"and",
	"in",
	"-",
	"that",
	"have",
}

func removeSpecialCharacters(token string) string {
	switch token[0] {
		case '(':
			return fmt.Sprintf("%s", token[1:len(token)])
	}
	switch token[len(token)-1] {
		case ')', '?', '.', ':', ',':
			return fmt.Sprintf("%s", token[:len(token)-1])

	}
	switch token[len(token)-1] {
		case '-':
			return ""
	}
	return token
}