package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ashish336b/wordle-cli/helper"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var round = 1

func play(solutionWord string) {
	for {
		var guess string
		fmt.Printf("Enter your guess %d/%d: ", round, helper.MAX_CHANCES)
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
	helper.ShowInstruction()
	play(helper.RandomSolution())
	return nil
}

func officialWordle(c *cli.Context) error {
	helper.ShowInstruction()
	play(helper.OfficialSolution())
	return nil
}

func playInfinity(c *cli.Context) error {
	helper.MAX_CHANCES = 50
	helper.ShowInstruction()
	play(helper.RandomSolution())
	return nil
}

func main() {
	err := (&cli.App{
		Name:   "wordle-go",
		Usage:  "wordle-in-terminal",
		Action: officialWordle,
		Commands: []*cli.Command{
			{
				Name:    "random",
				Aliases: []string{"r"},
				Usage:   "play the game",
				Action:  playRandom,
			},
			{
				Name:    "infinity",
				Aliases: []string{"i"},
				Usage:   "Play infinity",
				Action:  playInfinity,
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
