package main

import "fmt"

func BasicAtoi(s string) int {
	result := 0
	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}
	return result
}

func main() {
	fmt.Println(BasicAtoi("12356789"))
	fmt.Println(BasicAtoi("43"))
	fmt.Println(BasicAtoi("1579"))
}
