package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 11341
	b := 9
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
