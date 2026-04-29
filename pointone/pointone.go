package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	a := 42
	fmt.Println(a)
	PointOne(&a)
	fmt.Println(a)
}
