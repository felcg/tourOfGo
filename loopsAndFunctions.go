package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	guess := 1.0
	lastGuess := guess

	for i := 0; i <= 10; i++ {
		guess -= (guess*guess - x) / (2 * guess)
		
		if(lastGuess == guess) {
			return guess
		}
	}

	return guess
}

func main() {
	fmt.Println(Sqrt(25))
}