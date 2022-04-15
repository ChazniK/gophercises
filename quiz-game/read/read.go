package read

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadArguements() string {

	args := os.Args[1:]
	var normalizedInput = ""

	if len(args) == 1 {
		normalizedInput += strings.ToLower(args[0])
	}

	input := getInput(normalizedInput)

	return input
}

func getInput(input string) string {

	defaultFile := "problems.csv"

	if input == "" {
		return defaultFile
	}

	if input == "-h" || input == "-help" {
		printHelpText()
		return input
	}

	isMatch, _ := regexp.MatchString("[a-zA-Z0-9]+.csv", input)

	if isMatch {
		return input
	} else {
		printInvalidMessage(input)
		return "invalid"
	}
}

func printHelpText() {
	helpText := "Usage of ./quiz: " +
		"\n  -csv string" +
		"\n\ta csv file in the format of \"question/answer\" (default \"problems.csv\")\n" +
		"\n  -limit int" +
		"\n\tthe time limit for the quiz in seconds (default 30)\n"

	fmt.Printf("%s", helpText)
}

func printInvalidMessage(fileName string) {
	fmt.Printf("Specified file: %s is invalid, please check help flag\n", fileName)
}
