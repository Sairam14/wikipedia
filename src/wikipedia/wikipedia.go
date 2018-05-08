package main

import (
	"os"
	"fmt"
        "math"
)

const (
	ANSWERS = "wikipedia/resources/answers.txt"
	QUESTIONS = "wikipedia/resources/questions.txt"
	PARAGRAPH = "wikipedia/resources/wikipedia.txt"
	SEMICOLON = ";"
)

type Score struct {
	question Question
	answerScoreMap []map[string]Weightage
	preciseAnswer string
	weightageConsidered *Weightage
}

type Weightage struct {
	minimumOffsetDiffWeightage float64
	entityAnswerOccuranceWeightage int
	weightedAverage float64
}

func main(){
	paragraph := new(Paragraph)
	file, err := os.Open(PARAGRAPH)
	if err != nil {
		fmt.Println(err)
	}

	parseParagraph(file, 0, paragraph)
	fmt.Println(paragraph.text + "\n")

	questions,err := parseQuestions()
	if err != nil {
		fmt.Println(err)
	}

	answers, err :=  parseAnswers()
	if err != nil {
		fmt.Println(err)
	}

	scores := make([]Score, 0)

	for _,question := range questions {
		score := new(Score)
		score.question = question
		for _, answer := range answers {
			scoreAnswersForQuestion(paragraph, question, answer, score)
		}
		scores = append(scores, *score)
	}

	for _,score := range scores {
		evaluateAnswerByWeightedAverage(&score)
		//fmt.Println(score.question.text, "\n\"",score.preciseAnswer,"\"\n", "Weighted Average (Percentage) :: ",score.weightageConsidered.weightedAverage)
		fmt.Println(score.question.text, "\n",score.preciseAnswer,"\n")
	}
}

func evaluateAnswerByWeightedAverage(score *Score){
	score.preciseAnswer = ""
	score.weightageConsidered = new(Weightage)
	score.weightageConsidered.minimumOffsetDiffWeightage = math.MaxInt32
	score.weightageConsidered.entityAnswerOccuranceWeightage = -1
	score.weightageConsidered.weightedAverage = 0
	for _, result := range score.answerScoreMap {
		for key, weightage := range result {
			weightage.weightedAverage = getWeightedAverage(&weightage)
			if (score.weightageConsidered.weightedAverage < weightage.weightedAverage){
				score.weightageConsidered  = &weightage
				score.preciseAnswer = key
			}
		}
	}
}


func getWeightedAverage(weightage *Weightage)(float64){
	percent := (
		(((float64(100 - weightage.minimumOffsetDiffWeightage) * float64(100))) + // less the offset higher the weightage
		((float64(weightage.entityAnswerOccuranceWeightage) * float64(100)))) / // more the occurance of question tokens in answer. higher the weightage
			float64(100))
	if (percent > 100){
		return 100
	}
	return percent
}