package main

import "fmt"

func main() {

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condición
	value := 34
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es negativo")
	default:
		fmt.Println("No condition")
	}
}
