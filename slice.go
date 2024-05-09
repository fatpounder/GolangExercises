package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

//a slice does not store data, but rather describes a section of the referenced array
//when slicing, you can omit the high or low bounds to use their defaults instead
//default for low bound is zero and default for high bound is the length of the array
//for the array var a [10]int, the following slices are the same: a[0:10], a[:10], a[0:], a[:]
//the capacity of a slice is the number of elements in the referenced array, counting the first element in the slice
//the length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)
//you can extend a slice's length by re-slicing it, provided it has sufficient capacity
