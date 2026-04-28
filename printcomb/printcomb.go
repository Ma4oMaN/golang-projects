package main

import "fmt"

func PrintComb() {
	first := true

	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {

				if !first {
					fmt.Print(",")
				}
				fmt.Printf("%d%d%d", a, b, c)
				first = false
			}
		}
	}
	fmt.Println()
}

func main() {
	PrintComb()
}
