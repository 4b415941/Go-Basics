package main

import "fmt"

func main() {

	switch grade := 3; grade {
	case 1:
		fmt.Println("Failed")
		fallthrough
	case 2:
		fmt.Println("Failed")
		fallthrough
	case 3:
		fmt.Println("Passed")
	default:
		fmt.Println("Invalid")
	}
}
