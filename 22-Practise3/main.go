/* package main

import "fmt"

func main() {

	myArr := [5]string{"a", "b", "c", "d", "e"}

	for index, value := range myArr {
		fmt.Println(index, value)
	}

	fmt.Println()

	mySlc := []int{1, 2, 3, 4, 5}

	for i, v := range mySlc {
		fmt.Println(i, v)
	}

} */

package main

import "fmt"

func main() {

	myMap := map[string][]string{
		"Assasins": {"Akali", "Evelynn", "Kha'zix"},
		"Marksmen": {"Akshan", "Aphelios", "Jhin"},
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
