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
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input you guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value", err)
			continue
		}
		fmt.Println("Your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Answer is smaller. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Answer is bigger. Please try again")
		} else {
			fmt.Println("Correct!")
			break
		}
	}

}
