package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//range iterates over a slice or map
//when ranging over a slice, two values are returned for each iteration: first the index, then the copy of the element at that index
//can skip the index or value by assigning to _ ex:for _, value := range pow
//if you want only the index, you can omit the second variable
