package main

import "fmt"

func main() {
	/* 	var (
		name               = "Kadir"
		age                = 27
		isMarried          = false
		experience         = 2
		salary             = 500
		weight     float64 = 83.1
	) */

	/* var name, age, isMarried, experience, salary, weight = "Kadir", 27, false, 2, 500, 83.1 */
	/* name, age, isMarried, experience, salary, weight := "Kadir", 27, false, 2, 500, 83.1 */

	var name string
	var age int
	var trueorfalse bool
	/* fmt.Println(name, age, isMarried, experience, salary, weight) */
	fmt.Println(name)        //Zero Values stirng -> ""
	fmt.Println(age)         // Zero Values numeric -> 0
	fmt.Println(trueorfalse) // Zero Values bool -> False
}
