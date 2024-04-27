package main

import "fmt"

func main() {

	const x = 3 // typless
	var y int16 = 12

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)

	fmt.Printf("%T %v\n", x+y, x+y) // type conversion
}
