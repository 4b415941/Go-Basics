package main

import "fmt"

func main() {

	/* x := 33

	if x%2 == 0 {
		fmt.Println(x, "Odd")
	} else {
		fmt.Println(x, "Even")
	} */

	/* 	x := 12

	   	if x > 15 {
	   		fmt.Println("X is higher than 15")

	   	} else if x < 15 && x > 10 {
	   		fmt.Println("X between 10 and 15")

	   	} else {
	   		fmt.Println("X lower than 10")
	   	} */

	x := 11

	if x%2 == 0 {
		fmt.Println(x, "Even")
		return
	}
	fmt.Println(x, "Odd")
}

// if <boolean expression> {code} else {code}
