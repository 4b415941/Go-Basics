/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your exam result: ")
	grade, _ := getGrade()

	var result string

	if grade >= 50 {
		result = "Pass"
	} else {
		result = "Failed"
	}
	fmt.Println(result)
}

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print(err)
	}
	return num, nil
}
*/

// Number guessing game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := numRand(1, 5)

	fmt.Println("-----Guess a number between 1-100 Game-----")

	reader := bufio.NewReader(os.Stdin)
	for attempts := 1; attempts <= 10; attempts++ {
		fmt.Print("Guess >> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if num == target {
			fmt.Println("You win! The number is:", target)
			return
		} else {
			fmt.Printf("Wrong guess..! You have %d attempts left.\n", 10-attempts)
		}
	}

	fmt.Println("Game over! The correct number was:", target)
}

func numRand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
