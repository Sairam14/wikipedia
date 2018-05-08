package main

import (
	"regexp"
	"log"
	"strings"
)

var pluralDictionary = map[string]string{
	"zebras": "zebra",
	"aims" : "aim",
}

var singularDictionary = map[string]string{
	"zebra": "zebras",
	"aim" : "aims",
}

var numberList = map[string]int {
	"one" : 1,
	"two" : 2,
	"three" : 3,
	"four" : 4,
	"five" : 5,
	"six" : 6,
	"seven" : 7,
	"eight" : 8,
	"nine" : 9,
	"ten" : 10,
}

type Token struct {
	kind TokenKind
	text string
}

type TokenKind int

const (
	SpecialCharacter = 1+ iota
	String
	PluralString
	Number
)

func tokenize(text string)([]Token){
	const BufferSize = 1
	//runes := []rune(text)
	tokens := make([] Token, BufferSize)
	pos := 0

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	textValue := strings.ToLower(reg.ReplaceAllString(text, ""))
	if pluralDictionary[textValue] != "" {
		tokens[pos].kind = PluralString
	} else if numberList[strings.ToLower(textValue)] > 0 {
		tokens[pos].kind = Number
	} else {
		tokens[pos].kind = String
	}
	tokens[pos].text = textValue
	pos++

	return tokens
}

func checkPluralAndSingularToken(qToken Token, answerToken Token) (bool, string){
	if (qToken.kind == PluralString && answerToken.kind == String){
		singularText := pluralDictionary[qToken.text]
		if singularText == ""{
			singularText = qToken.text
		}
		//fmt.Println(singularText, answerToken.text, strings.EqualFold(singularText, answerToken.text))
		return strings.EqualFold(singularText, answerToken.text), singularText
	}

	if (qToken.kind == String && answerToken.kind == PluralString) {
		pluralText := singularDictionary[qToken.text]
		if pluralText == ""{
			pluralText = qToken.text
		}
		//fmt.Println(pluralText, answerToken.text, strings.EqualFold(pluralText, answerToken.text))
		return strings.EqualFold(pluralText, answerToken.text), pluralText
	}

	//fmt.Println(qToken.text, answerToken.text, strings.EqualFold(qToken.text, answerToken.text))
	return strings.EqualFold(qToken.text, answerToken.text), qToken.text

}