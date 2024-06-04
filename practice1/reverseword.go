package main

import (
	"fmt"
	"strings"
)

func ReverseWithJoin(s string) string {
	chars := strings.Split(s, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return strings.Join(chars, "")
}

func main() {
	original := "greatbigword"
	reversed := ReverseWithJoin(original)
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", reversed)
}
