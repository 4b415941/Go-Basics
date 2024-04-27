package main

import "fmt"

func main() {

	var name = "Kadir"
	var age = 27
	var isMarried = false
	var experience uint8 = 2
	var salary uint16 = 500
	var weight float32 = 83.1

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", experience)
	fmt.Printf("%T\n", salary)
	fmt.Printf("%T\n", weight)
}
