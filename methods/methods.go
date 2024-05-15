package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

//no classes in Go, but can define methods on types
//a method is a function with a special "receiver" argument
//the receiver appears as its own argument list between the func keyword and the method name
//here, the Abs methods has a receiver of type Vertex named v
//Remember, methods ARE functions, just with a receiver argument
