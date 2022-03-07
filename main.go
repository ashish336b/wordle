package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ashish336b/wordle-in-terminal/helper"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func officialWordle(c *cli.Context) error {
	fmt.Println(helper.CheckAccuracy("hello", "hiiei"))
	return nil
}

func isGameOver(round int, solutionWord string, guessWord string) bool {
	if round == 6 && solutionWord != guessWord {
		return true
	}
	if solutionWord == guessWord {
		return true
	}
	return false
}

func playGame(c *cli.Context) error {
	solutionWord := helper.GetSolution()
	round := 1
	for {
		var guess string
		fmt.Printf("Enter your guess %d/%d: ", round, 6)
		fmt.Scanln(&guess)
		guess = strings.ToUpper(guess)
		result := helper.CheckAccuracy(solutionWord, guess)
		if guess == solutionWord {
			if round == 6 {
				color.Red("You lose!\n")
				fmt.Printf("correct word is: %s\n", solutionWord)
				break
			}
			fmt.Println(result)
			color.Green("You win!")
			break
		}
		fmt.Println(result)
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
				Action:  officialWordle,
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
