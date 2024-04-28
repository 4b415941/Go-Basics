package main

import (
	"fmt"
)

func main() {

	/* myMap := map[string]int{
		"Kadir":  26,
		"Can":    6,
		"random": 5,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Kadir"]) */

	/* isMarried := map[string]bool{
		"Kadir": false,
		"Can":   false,
		"James": true,
	}
	fmt.Println(isMarried) */

	studentGrades := map[string]int{
		"Kadir": 70,
		"Can":   50,
		"James": 80,
	}

	fmt.Println(studentGrades)
	studentGrades["Beng"] = 80
	fmt.Println(studentGrades)
	delete(studentGrades, "Kadir")
	fmt.Println(studentGrades)

}
