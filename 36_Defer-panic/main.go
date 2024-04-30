package main

import (
	"fmt"
)

func main() {
	/* defer fmt.Println("Hello")
	fmt.Println("World") */

	/* x := 10
	y := 20

	defer fmt.Println("x:", x)
	defer fmt.Println("y:", y)

	x = 100
	y = 200
	fmt.Println("x:", x)
	fmt.Println("y:", y) */

	var condition = true

	defer cleanup()
	if condition {
		panic("An error occurred!")
	}

}

func cleanup() {
	fmt.Println("Cleanup worked..")
}
