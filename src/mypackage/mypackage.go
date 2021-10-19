package mypackage

import "fmt"

// CarPublic: Car with public access
type CarPublic struct {
	Brand string
	year  int
}

//Car Private
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage() {
	fmt.Println("Hola")
}

// printMessage1 imprimir un mensaje
func printMessage1() {
	fmt.Println("Hola")
}
