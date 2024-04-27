package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func(Pack) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}
	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)
	/* fmt.Println(blockVar) */ // out scope

	myFunc()
}

func myFunc() {
	/* fmt.Println(packVar) */
	/* fmt.Println(funcVar) */ //out scope
	fmt.Println(funcVar)
}
