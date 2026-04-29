package main

import "fmt"

func PrintNbr(n int) {
	if n == 0 {
		fmt.Print("0")
	}

	if n < 0 {
		fmt.Print("-")
	}

	recursivePrint(n)
	fmt.Println()
}

func recursivePrint(n int) {
	if n == 0 {
		return
	}

	recursivePrint(n / 10)

	mod := n % 10

	if mod < 0 {
		mod = -mod
	}

	fmt.Printf("%c", mod+'0')
}

func main() {
	PrintNbr(0)
	PrintNbr(123)
	PrintNbr(-456)
}
