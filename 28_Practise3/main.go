/* package main

import "fmt"

type myType int

func main() {

	var x myType
	x = 10

	fmt.Printf("%T, %v", x, x)
}
*/

/* package main

import "fmt"

func main() {

	x := 10
	y := (&x)
	*y = 20

	fmt.Println(x)
}
*/

/* package main

import (
	"fmt"
)

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("rectangle side length: ", r.a, "-", r.b)
}

func (r rectangle) area() int {
	return r.a * r.b
}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)
}

func main() {
	r1 := rectangle{3, 5}
	r1.display()
	fmt.Println("Area:", r1.area())
	fmt.Println("Circumference:", r1.circumference())
}
*/

package main

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	f1 := family{
		name:      "Kadircan",
		age:       26,
		isMarried: false,
	}
	f2 := family{
		name:      "Beng",
		age:       24,
		isMarried: false,
	}
	f3 := family{
		name:      "Nur",
		age:       26,
		isMarried: false,
	}
	var f4 family
	f4.name = "Sally"
	f4.age = 21
	f4.isMarried = true

	return []family{f1, f2, f3, f4}
}
func main() {
	families := showFamily()

	for i := 0; i < len(families); i++ {
		fmt.Println(families[i])
	}
}
