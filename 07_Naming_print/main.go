package main

import "fmt"

func main() {

	/* fmt.Println("Merhaba") //newline
	fmt.Print("Merhaba")
	fmt.Printf("Merhaba") */

	name := "Kadir"

	/* fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name) */

	fmt.Print("My name is:", name)
	fmt.Println("")
	fmt.Println("My name is:", name)
	fmt.Printf("My name is: %v %T", name, name)
}
