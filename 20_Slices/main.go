package main

import "fmt"

func main() {

	/* mySlc := []int{1, 2, 3} //Literal method
	fmt.Println(mySlc)
	fmt.Print(len(mySlc)) */
	/*
		var mySlc []int
		mySlc = make([]int, 4)	// Make Method
		fmt.Println(mySlc) */

	/* mySlc := []int{1, 2, 3}
	fmt.Println(mySlc)
	mySlc2 := mySlc
	fmt.Println(mySlc2)

	mySlc2[0] = 100
	fmt.Println(mySlc2)
	fmt.Println(mySlc) // pass by reference

	myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)
	myArr2 := myArr
	fmt.Println(myArr2)

	myArr2[0] = 100
	fmt.Println(myArr2)
	fmt.Println(myArr) // pass by value */

	/* underArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(underArr)

	mySlc := underArr[2:5]
	fmt.Println(mySlc) */

	/* mySlc := []int{1, 2, 3}
	mySlc2 := []int{4, 5}

	mySlc = append(mySlc, mySlc2...)
	fmt.Println(mySlc) */

	var mySlc []int
	fmt.Printf("%#v", mySlc) // []int(nil)

	fmt.Println()

	mySlc2 := make([]int, 0)
	fmt.Printf("%#v", mySlc2) // []int{}

}
