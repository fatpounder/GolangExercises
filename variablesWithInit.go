package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//a var declaration can include initializers, one per variable
//if an initializer is present, the type can be omitted and the variable will take the type of the initializer
