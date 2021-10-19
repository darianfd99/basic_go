package main

import (
	"fmt"
	pk "go_basico/src/mypackage"
)

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"

	fmt.Println(otherCar)

	// Public
	var CarPublica pk.CarPublic
	CarPublica.Brand = "Ferrari"

	fmt.Println(CarPublica)

	//Private
	//var myOtherCar pk.carPrivate
	//fmt.Println(myOtherCar)

	pk.PrintMessage()
	pk.printMessage1()
}
