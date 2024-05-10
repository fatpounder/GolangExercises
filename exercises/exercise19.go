package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\t", x)

	//if x <= 100 {
	//	fmt.Println("x is less than 100")
	//} else if x >= 101 && x <= 200 {
	//	fmt.Println("x is between 101 and 200")
	//} else if x >= 201 && x <= 250 {
	//	fmt.Println("x is between 201 and 250")
	//} else {
	//	fmt.Println("x is more than 250")
	//}

	switch {
	case x <= 100:
		fmt.Println("x is less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("x is between 201 and 250")
	default:
		fmt.Println("x is more than 250")
	}
}
