package main

import (
	"quiz/process"
	"quiz/read"
)

func main() {

	input := read.ReadArguements()

	if input == "-h" || input == "-help" || input == "invalid" {
		return
	}

	questionsAndAnswers := process.ProccessFile(input)

	for i := 0; i < len(questionsAndAnswers); i++ {
	}
}
