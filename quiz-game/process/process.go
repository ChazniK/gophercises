package process

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ProccessFile(fileName string) []questionAnswer {

	var questionsAndAnswers []questionAnswer

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questionsAndAnswers = append(questionsAndAnswers, createQuestionAnswer(record))
	}

	return questionsAndAnswers
}

func createQuestionAnswer(record []string) questionAnswer {
	var newQuestionAnswer questionAnswer
	newQuestionAnswer.question = record[0]
	newQuestionAnswer.answer = record[len(record)-1]

	return newQuestionAnswer
}

type questionAnswer struct {
	question string
	answer   string
}
