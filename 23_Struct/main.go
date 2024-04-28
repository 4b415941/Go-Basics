/*
	package main

import "fmt"

func main() {

		var employee struct {
			name      string
			age       int
			isMarried bool
		}

		employee.name = "Kadircan"
		employee.age = 26
		employee.isMarried = false

		fmt.Println(employee)

		var employee2 struct {
			name      string
			age       int
			isMarried bool
		}

		employee2.name = "Beng"
		employee2.age = 24
		employee2.isMarried = false

		fmt.Println(employee2)
	}
*/
package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
	kids      []string
}

func main() {

	var e1 employee
	e1.name = "Kadircan"
	e1.age = 26
	e1.isMarried = false

	var e2 employee
	e2.name = "Beng"
	e2.age = 24
	e2.isMarried = false

	e3 := employee{
		name:      "Random",
		age:       20,
		isMarried: true,
	}
	fmt.Println(e1, e2, e3)

	e4 := employee{
		name:      "Sally",
		age:       50,
		isMarried: true,
		kids: []string{
			"Kadircan",
			"Nur",
		},
	}
	fmt.Println(e4.kids)
}
