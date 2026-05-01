package main

import "fmt"

func PrintComb2() {
	first := true

	for ab := 0; ab <= 99; ab++ {
		for cd := ab + 1; cd <= 99; cd++ {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%02d %02d", ab, cd)
			first = false
		}
	}
	fmt.Println()
}

func main() {
	PrintComb2()
}
