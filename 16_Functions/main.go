/* package main

import (
	"fmt"
)

func main() {

	/* fmt.Println(sum(5, 10))
	hello()
	hello("Kadircan", 27) //args
}

/* func sum(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello Func")
}

func hello(name string, age int) { //parametre
	fmt.Printf("name:%s age:%d", name, age)
}
*/

/* package main

import (
	"fmt"
	"time"
)

func main() {

	var x int = 10
	var today time.Time = time.Now() // Now -> method

	fmt.Println(x, today)
} */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your exam result: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ blank identifier
	fmt.Println(value)
}
