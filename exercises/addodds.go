package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 27; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of odd numbers from 1 to 27: %d\n", sum)
}
