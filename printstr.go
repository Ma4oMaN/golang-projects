package main

import "fmt"

func PrintStr(s string) {
	for _, c := range s {
		fmt.Printf("%c", c)
	}
}

func main() {
	PrintStr("Hello, Мир! 😃")
	fmt.Println()
}
