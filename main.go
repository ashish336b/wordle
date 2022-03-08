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

var round = 1

func play(solutionWord string) {
	for {
		var guess string
		fmt.Printf("Enter your guess %d/%d: ", round, 6)
		fmt.Scanln(&guess)
		guess = strings.ToUpper(guess)
		message, isValid := helper.ValidateWord(guess)
		if !isValid {
			color.Red(message)
			play(solutionWord)
			break
		}
		hint := helper.CheckAccuracy(solutionWord, guess)
		if helper.WinCondition(round, solutionWord, guess) {
			fmt.Println(hint)
			color.Green("you Win")
			break
		}
		if helper.LoseCondition(round, solutionWord, guess) {
			color.Red("you lost!")
			fmt.Printf("the word you are looking for is %s \n", helper.GetRightAnswer(solutionWord))
			break
		}
		fmt.Println(hint)
		round++
		play(solutionWord)
		break

	}
}

func playRandom(c *cli.Context) error {
	solutionWord := helper.RandomSolution()
	play(solutionWord)
	return nil
}

func officialWordle(c *cli.Context) error {
	play(helper.OfficialSolution())
	return nil
}

func main() {
	err := (&cli.App{
		Name:   "wordle-go",
		Usage:  "wordle-in-terminal",
		Action: playRandom,
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
