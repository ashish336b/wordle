package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ashish336b/wordle-in-terminal/helper"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func playGame(c *cli.Context) error {
	rand.Seed(time.Now().UnixNano())
	solutionIndex := rand.Intn(len(helper.Words))
	solutionWord := strings.ToUpper(helper.Words[solutionIndex])
	solutionWordChars := strings.Split(solutionWord, "")
	fmt.Println(solutionWordChars)
	i := 1
	for {
		var guess string
		fmt.Printf("Enter your guess %d/%d: \n", i, 6)
		fmt.Scanln(&guess)
		guess = strings.ToUpper(guess)
		if guess == solutionWord || i == 6 {
			if i == 6 {
				fmt.Println("you lose!")
				break
			}
			color.Green("You win!")
			break
		}
		i++
	}
	return nil
}
func main() {
	err := (&cli.App{
		Name:   "wordle-go",
		Usage:  "wordle-in-terminal",
		Action: playGame,
		Commands: []*cli.Command{
			{
				Name:    "play",
				Aliases: []string{"p"},
				Usage:   "play the game",
				Action: func(c *cli.Context) error {
					fmt.Println("hello world")
					return nil
				},
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
