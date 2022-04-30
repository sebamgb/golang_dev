package main

import (
	"fmt"
	"strings"
)

func isPalindrom(text string) {
	text = strings.ToLower(text)
	var text_revers string
	for i := len(text) - 1; i >= 0; i-- {

		text_revers += string(text[i])

	}
	if text == text_revers {
		fmt.Println("es palindromo")
	} else {
		fmt.Println("no es palindromo")
	}
}

func main() {
	// Range
	slice := []string{"hola", "que", "hace"}
	for i, value := range slice {
		fmt.Println(i, value)
	}
	// escape de indice
	for _, value := range slice {
		fmt.Println(value)
	}
	//escape valor
	for i := range slice {
		fmt.Println(i)
	}

	// Palindromo?
	isPalindrom("Ama")
}
