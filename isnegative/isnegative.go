package main

import "fmt"

func IsNegative(nb int) bool {
	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Print("F")
	}
}

func main() {
	IsNegative(7)
	IsNegative(-3)
}
