package main

import "fmt"

type shape interface {
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
	}
}

type triangle struct {
	a, h float64
}

type square struct {
	a float64
}

type rectangle struct {
	a, b float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

func (s square) area() float64 {
	return (s.a * s.a)
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func main() {

	t1 := triangle{3, 5}
	s1 := square{5}
	r1 := rectangle{5, 8}

	printArea(t1, s1, r1)
}
