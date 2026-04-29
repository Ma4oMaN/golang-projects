package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	mod := *a % *b
	*a = div
	*b = mod
}

func main() {
	c := 15643
	d := 423
	UltimateDivMod(&c, &d)
	fmt.Println(c)
	fmt.Println(d)
}
