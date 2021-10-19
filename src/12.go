package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB de disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 15, brand: "msi", disk: 100}
	fmt.Println(myPC)
}
