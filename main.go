package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ashish336b/wordle-in-terminal/helper"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	solutionIndex := rand.Intn(len(helper.Words))
	solutionWord := strings.ToUpper(helper.Words[solutionIndex])
	solutionWordChars := []rune(solutionWord)
	fmt.Println(solutionWordChars)
	for {
		var guess string
		fmt.Scanln(&guess)
		guess = strings.ToUpper(guess)
		if guess == solutionWord {
			fmt.Println("You win!")
			break
		}
		fmt.Println("Nope, try again")
	}
}
