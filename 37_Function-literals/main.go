package main

import (
	"fmt"
)

func main() {

	/* func() {
		fmt.Println("f1")
	}() */
	/*
		func(x, y int) {
			fmt.Println(x + y)
		}(3, 5) */

	add := func(x, y int) int {
		return x + y
	}

	multiply := func(x, y int) int {
		return x * y
	}

	/* fmt.Println(reflect.TypeOf(add))
	fmt.Println(add(3, 10)) */

	a, b := calculator(3, 5, add, multiply)
	fmt.Println(a, b)

}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
