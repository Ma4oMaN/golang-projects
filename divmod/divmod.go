package main

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	var c, d int
	DivMod(235, 78, &c, &d)
	fmt.Println(c)
	fmt.Println(d)
}
