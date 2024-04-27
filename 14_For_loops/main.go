package main

import "fmt"

func main() {

	// for <init statement>; <condition>; <post statement>

	/* 	for i := 1; i < 10; i++ {
		fmt.Println(i)
	} */

	/* for {
		fmt.Println("Infinite Loop")
	} */

	i := 0

	for ; i <= 10; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		} else {
			fmt.Println(i)
		}
	}

}
