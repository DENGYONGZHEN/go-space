package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessages(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func exercise() {

	msg = "Hello,world!"

	wg.Add(1)
	go updateMessages("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessages("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessages("Hello,world!")
	wg.Wait()
	printMessage()
}
