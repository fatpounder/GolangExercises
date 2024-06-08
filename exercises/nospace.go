package main

import (
	"fmt"
	"strings"
)

func removeSpaces(sentence string) string {
	return strings.ReplaceAll(sentence, " ", "")
}

func main() {
	input := "This is a sentence"
	result := removeSpaces(input)
	fmt.Println("Result:", result)
}
