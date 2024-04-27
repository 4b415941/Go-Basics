package main

import "fmt"

/* Comparison Operators
== Equal != not Equal
< less than > greater than
<= less than or eq >= greater than or eq

Logical Operators

&& conditional AND	p && q is "if p than q else false"
|| conditional OR	p || q is "if p than true else q"
! NOT 				!p     is "not p" */

func main() {

	/* x, y := 5, 9

	fmt.Printf("%T, %v\n", x == y, x == y)
	fmt.Printf("%T, %v\n", x != y, x != y)
	fmt.Printf("%T, %v\n", x < y, x < y)
	fmt.Printf("%T, %v\n", x > y, x > y)
	fmt.Printf("%T, %v\n", x >= y, x >= y)
	fmt.Printf("%T, %v\n", x <= y, x <= y) */

	x, y := 15, 20

	set1 := (x == y) // false
	set2 := (x <= y) // true

	fmt.Printf("%T, %v\n", set1, set1)
	fmt.Printf("%T, %v\n", set2, set2)

	fmt.Printf("%v\n", (set1 && set2))
	fmt.Printf("%v\n", (set1 || set2))

}
