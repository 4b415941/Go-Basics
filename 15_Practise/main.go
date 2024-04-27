package main

import "fmt"

func main() {

	var i, j int
	for i = 2; i <= 50; i++ {
		isPrime := true
		for j = 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
