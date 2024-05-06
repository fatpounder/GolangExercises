package main

import "fmt"

const (
	_ = iota
	g
	h
	m
	n
	k
	l
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<g, 1<<g)
	fmt.Printf("%d \t %b\n", 1<<h, 1<<h)
	fmt.Printf("%d \t %b\n", 1<<m, 1<<m)
	fmt.Printf("%d \t %b\n", 1<<n, 1<<n)
	fmt.Printf("%d \t %b\n", 1<<k, 1<<k)
	fmt.Printf("%d \t %b\n", 1<<l, 1<<l)
}
