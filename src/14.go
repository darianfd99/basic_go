package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(2)
	go say("World", &wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Adios")

	}(&wg)

	wg.Wait()

	fmt.Println("Allllo")
	//time.Sleep(time.Second * 1)
}
