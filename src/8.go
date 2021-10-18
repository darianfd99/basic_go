package main

import (
	"fmt"
	"strings"
)

func isPalidromo(text string) {
	text = strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es un palíndromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for _, value := range slice {
		fmt.Println(value)
	}

	isPalidromo("Ana")
}
