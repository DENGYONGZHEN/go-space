package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	//🎈 third step
	defer wg.Done()
	fmt.Println(s)
}
func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	//🎈 first step
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	//🎈 second step,wait count to be 0
	wg.Wait()

	////🎈 third step
	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
