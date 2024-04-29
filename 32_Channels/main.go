/* package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) display() {
	fmt.Println("A circle")
	//wg.Done()
}

//var wg sync.WaitGroup

func main() {

	c1 := circle{5}
	//wg.Add(1)
	go c1.display()
	//wg.Wait()
	area1 := c1.area()
	fmt.Printf("%.2f\n", area1)
}
*/

/* package main

import "fmt"

func hello(ch1 chan string) {
	ch1 <- "Hello my friend!"
}

func main() {
	ch := make(chan string)
	go hello(ch)
	fmt.Println(<-ch)
}
*/

/* package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) display() {
	fmt.Println("A circle..")
}

func main() {

	c1 := circle{7}
	c1.display()
	area1 := go c1.area() // you cant use with return.
	fmt.Println(area1)
}
*/

package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func area(c circle, myChannel chan float64) {
	result := math.Pi * c.r * c.r
	myChannel <- result
}

func (c circle) display() {
	fmt.Println("A circle..")
}

func main() {

	c1 := circle{7}
	ch := make(chan float64)

	go area(c1, ch)
	fmt.Printf("%.2f", <-ch)
	fmt.Println()
	c1.display()
}
