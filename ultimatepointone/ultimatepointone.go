package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 1234
	b := &a
	c := &b
	d := &c
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	UltimatePointOne(d)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
