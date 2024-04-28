package main

import "fmt"

func main() {

	//name := "Kadircan"
	/*   	fmt.Println(name)
	fmt.Println(&name)
	*/
	// x := 22
	/* 	y := &x    // *int
	   	z := &name // *string

	   	fmt.Printf("%T, %v", x, x)
	   	fmt.Println()
	   	fmt.Printf("%T, %v", y, y)
	   	fmt.Println()
	   	fmt.Printf("%T, %v", z, z) */

	/* 	fmt.Println(x)     //22
	   	fmt.Println(&x)    //address
	   	fmt.Println(*(&x)) //dereferencing */

	/* x1 := 10
	x2 := &x1

	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	*x2 = 3
	fmt.Println(x1, *x2) */

	//x1 := [4]int{1, 10, 100, 1000} // array pass by value
	x1 := []int{1, 10, 100, 1000} //slice pass by reference
	x2 := x1

	fmt.Println(x1, x2)
	x2[0] = 3
	fmt.Println(x1, x2)
}
