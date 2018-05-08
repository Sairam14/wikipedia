package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

type Answer struct {
	text string
	tokens []Token
}

func parseAnswers() ([]Answer, error){
	file, err := os.Open(ANSWERS)
	answers := []Answer{}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}

		result := strings.Split(scanner.Text(), SEMICOLON);
		for i := range result {
			answer := new(Answer)
			answer.text = strings.ToLower(result[i])
			words := strings.Fields(answer.text)
			exists :=  false
			for _,word := range words {
				exists = false
				actualWord := removeSpecialCharacters(word)
				for _, ignoreToken := range ignoreTokens {
					if strings.EqualFold(actualWord, ignoreToken) {
						exists = true
						break
					}
				}
				if !exists {
					answer.tokens = append(answer.tokens, tokenize(actualWord)...);
				}
			}
			answers = append(answers, *answer)
		}
	}
	//fmt.Println(answers)
	return answers, nil
}

