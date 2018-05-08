package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

type Question struct {
	text string
	tokens []Token
}

func parseQuestions() ([]Question, error){
	const BufferSize = 1
	questions := []Question{}
	file, err := os.Open(QUESTIONS)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			break
		}
		line := scanner.Text()
		question := new(Question)
		question.text = line
		words := strings.Fields(line)
		exists :=  false
		for _, word := range words {
			exists = false
			actualWord := removeSpecialCharacters(word)
			for _, ignoreToken := range ignoreTokens {
				if strings.EqualFold(actualWord, ignoreToken) {
					exists = true
					break
				}
			}
			if !exists {
				question.tokens = append(question.tokens, tokenize(actualWord)...)
			}
		}

		questions = append(questions, *question)
	}
	//fmt.Println(questions)
	return questions, err
}

