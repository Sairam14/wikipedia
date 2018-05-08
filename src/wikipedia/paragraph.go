package main

import (
	"bufio"
	"strings"
	"io"
	"math"
)

type Paragraph struct {
	text string
	wordToTokens map[string][]Token
	wordToOffsets map[string][]int
}

func parseParagraph(input io.ReadSeeker, start int, paragraph *Paragraph) error {
	if _, err := input.Seek((int64)(start), 0); err != nil {
		return err
	}

	scanner := bufio.NewScanner(input)
	paragraph.wordToTokens = make(map[string][]Token)
	paragraph.wordToOffsets = make(map[string][]int)

	pos := start
	scanWords := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		start = pos
		pos += advance
		return
	}
	scanner.Split(scanWords)

	for scanner.Scan() {
		word := scanner.Text()
		exists := false
		for _, ignoreToken := range ignoreTokens {
			if strings.EqualFold(word, ignoreToken) {
				exists = true
				break
			}
		}
		if (!exists) {
			keyword := removeSpecialCharacters(strings.ToLower(word));
			evaluateTokens(paragraph, keyword)
			evaluateOffset(paragraph, start, keyword)
		}
	}
	setSentenceText(input, 0, paragraph)
	//fmt.Println(paragraph.wordToTokens)
	//fmt.Println(paragraph.wordToOffsets)
	return scanner.Err()
}

func answerWithMinimalOffSetDifference(paragraph *Paragraph, questionToken string, answerText string)(float64){
	originalOffsetOfAnswer := strings.LastIndex(paragraph.text, answerText)

	lastMinimalOffsetDifference := math.MaxFloat64
	for _, offset := range paragraph.wordToOffsets[questionToken]{
		if (offset == 0){
			continue
		}
		minimalOffsetDifference := 0.0
		if (originalOffsetOfAnswer > offset){
			//fmt.Println("answer offset in paragraph", originalOffsetOfAnswer, offset)
			minimalOffsetDifference = math.Abs((float64)(originalOffsetOfAnswer - offset))
		} else {
			//fmt.Println("answer offset in paragraph", originalOffsetOfAnswer, offset)
			minimalOffsetDifference = math.Abs((float64)(offset - originalOffsetOfAnswer))
		}
		if (lastMinimalOffsetDifference > minimalOffsetDifference){
			lastMinimalOffsetDifference = minimalOffsetDifference
		}

		//3fmt.Println("Max Value", minimalOffsetDifference, lastMinimalOffsetDifference, questionToken, answerText)
	}
	return lastMinimalOffsetDifference
}

func scoreAnswersForQuestion(sentence *Paragraph, question Question, answer Answer, score *Score) {
	scoreMap := (map[string]Weightage)(nil)
	weightage := new(Weightage)
	minimalDiffernceEvaluated := false
	finalOffSetValue := math.MaxFloat64
	minimalOffsetDiffBetweenQuestionAndAnswer := 0.0
	anyNumberTokenDetectedInQuestion := false
	for _, qToken := range question.tokens {
		scoreCounter := 0
		//fmt.Println(qToken.text)
		if !anyNumberTokenDetectedInQuestion && qToken.kind == Number {
			anyNumberTokenDetectedInQuestion = true
		}
		minimalDiffernceEvaluated = false
		for _, answerToken := range answer.tokens {
			matched, consideredText := checkPluralAndSingularToken(qToken, answerToken)
			if (matched){
				if anyNumberTokenDetectedInQuestion {
					scoreCounter++
				}
				scoreCounter++
				weightage.entityAnswerOccuranceWeightage = scoreCounter
			}

			if minimalDiffernceEvaluated == false {
				minimalOffsetDiffBetweenQuestionAndAnswer =
					answerWithMinimalOffSetDifference(sentence, consideredText, answer.text)
				if finalOffSetValue > minimalOffsetDiffBetweenQuestionAndAnswer {
					finalOffSetValue = minimalOffsetDiffBetweenQuestionAndAnswer
				}
				//fmt.Println(consideredText, answer.text, "OffSet returned", finalOffSetValue)
				minimalDiffernceEvaluated = true
			}
		}
	}
	//fmt.Println(question.text, answer.text, "Min Offset returned", finalOffSetValue)
	weightage.minimumOffsetDiffWeightage = finalOffSetValue
	if (scoreMap == nil) {
		scoreMap = make(map[string]Weightage)
	}
	scoreMap[answer.text] = *weightage

	if len(scoreMap) > 0 {
		score.answerScoreMap = append(score.answerScoreMap, scoreMap)
	}
}

func evaluateTokens(paragraph *Paragraph, key string){
	if paragraph.wordToTokens[key] == nil {
		paragraph.wordToTokens[key] = tokenize(key)
	} else {
		paragraph.wordToTokens[key] = append(paragraph.wordToTokens[key], tokenize(key)...)
	}
}

func evaluateOffset(paragraph *Paragraph, offset int, key string){
	if paragraph.wordToOffsets[key] == nil {
		offsets := make([]int, 1)
		paragraph.wordToOffsets[key] = append(offsets, offset)
	} else {
		paragraph.wordToOffsets[key] = append(paragraph.wordToOffsets[key], offset)
	}
}

func setSentenceText(input io.ReadSeeker, start int, sentence *Paragraph) (error) {
	if _, err := input.Seek((int64)(start), 0); err != nil {
		return err
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		sentence.text = strings.ToLower(scanner.Text())
	}
	return nil
}
