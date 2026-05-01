package main

import "fmt"

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	negative := false
	start := 0

	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		negative = true
		start = 1
	}

	for _, char := range s[start:] {
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	if negative {
		return -result
	}
	return result
}

func main() {
	fmt.Println(Atoi("13zxc-"))
	fmt.Println(Atoi("+13412"))
	fmt.Println(Atoi("-1214"))
	fmt.Println(Atoi("13-"))
	fmt.Println(Atoi("+13412"))
	fmt.Println(Atoi("fdhg76"))
}
