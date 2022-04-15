package main

import (
	"quiz/read"
)

func main() {

	input := read.ReadArguements()

	if input == "-h" || input == "-help" || input == "invalid" {
		return
	}

	println(input)
}
