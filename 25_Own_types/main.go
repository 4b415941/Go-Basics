/* package main

import (
	"fmt"
	"strings"
) */

// struct -> underlying type, employee -> defined type or named type

/* type mile float64
type kilometer float64
type myString string

func main() { */

/* 	f1 := float64(4.4)
fmt.Println()
fmt.Println(m1 + mile(f1)) */

// var m1 mile
// m1 = 3.2
// fmt.Println(m1)
// fmt.Printf("%T %v", m1, m1)

/* 	k1 := kilometer(7.2)
   	fmt.Println(k1)
   	fmt.Printf("%T %v", k1, k1)
	fmt.Println(m1 + k1) */

/* m2 := mile(2.5)

fmt.Println(m1 + m2) */
/*
	mystring := myString("Kadircan")
	fmt.Println(strings.ToUpper(string(mystring)))
} */

package main

import "fmt"

// struct -> underlying type, employee -> defined type or named type

type mile float64
type kilometer float64

func main() {
	m1 := mile(10)
	k1 := toKilometer(m1)
	fmt.Println(k1)
}

func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}
