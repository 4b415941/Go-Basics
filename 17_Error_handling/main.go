/* package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("OK, Its even : ", result)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("The number you entered is odd!")
	}
	return num, nil
}
*/

/* package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := sRoot(9)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(result)
	}

}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("please enter positive number")
	}
	return math.Sqrt(num), nil
}
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("File:", file)
	}
}
