package main

import "fmt"

func main() {
	value := "katak"
	konsonan := "a"

	fmt.Println(palindrome(value))
	fmt.Println(konsonanFunc(konsonan))
}

func palindrome(value string) bool {
	chars := []rune(value)
	var result string

	number := len(chars) - 1

	for i := 0; i < len(chars); i++ {
		result += string(chars[number-i])
	}

	if value == result {
		return true
	}

	return false
}

func konsonanFunc(value string) string {
	hurufVokal := []string{"a", "i", "u", "e", "o"}

	for _, huruf := range hurufVokal {
		if value == huruf {
			return "Vokal"
		}
	}

	return "Konsonan"
}
