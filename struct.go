package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//func main() {
//	fmt.Println(Vertex{1, 2})
//}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

//a struct is a collection of fields
//struct fields are accessed using a dot
