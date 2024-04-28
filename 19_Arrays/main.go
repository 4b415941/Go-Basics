package main

import "fmt"

func main() {

	// cities := [4]string{"Istanbul", "Belgrad", "Amsterdam", "Paris"}
	// cities := []string{"Istanbul", "Belgrad", "Amsterdam", "Paris"}
	// cities := [...]string{"Istanbul", "Belgrad", "Amsterdam", "Paris"}

	/* fmt.Println(cities)
	fmt.Println(cities[1])
	fmt.Println(len(cities))
	*/

	/* for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	} */

	/* 	for index, city := range cities {
		fmt.Println(index, city)
	} */

	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myArr = mySquare(myArr)
	fmt.Println(myArr)
}

func mySquare(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
