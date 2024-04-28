package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {

	/* e1 := employee{
		name:      "Kadircan",
		age:       27,
		isMarried: false,
	}
	fmt.Println(e1) */
	/* 	e2 := e1
	   	fmt.Println(e2)
	   	e2.name = "Beng"
	   	fmt.Println(e2)
	   	fmt.Println(e1) */

	m1 := manager{
		employee: employee{
			name:      "Beng",
			age:       20,
			isMarried: false,
		},
		hasDegree: false,
	}

	m2 := manager{}
	m2.name = "Kadircan"
	m2.age = 26
	m2.isMarried = false
	m2.hasDegree = true
	fmt.Println(m1)
	fmt.Println(m2)

	// Anonim Struct
	theBoss := struct {
		name  string
		money bool
	}{name: "James", money: true}

	fmt.Println(theBoss)
}
