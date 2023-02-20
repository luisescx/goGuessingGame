package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	secret := getRandomNumber()
	
	for matching := false; !matching; {
		guess := getUserInput()
		matching = isMatching(secret, guess)
	}
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

func getUserInput() int {
	fmt.Print("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed: ", input)
	}

	return input
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Too Big")
		return false
	} else if guess < secret {
		fmt.Println("Too Small")
		return false
	} else {
		fmt.Println("You Win!")
		return true
	}
}