package main

import "fmt"

func main() {
	myChannel := make(chan string)
	done := make(chan bool)
	/*
		go func() {
			myChannel <- "Hello World"
		}()
		// Blocking
		message, isOpen := <-myChannel
		fmt.Println(message, isOpen) */

	go func() {
		message := <-myChannel
		fmt.Println(message)
		done <- true
	}()
	go func() {
		myChannel <- "Hello world!"
	}()
	<-done
	fmt.Println("End of the main")
}
