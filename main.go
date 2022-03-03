package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ashish336b/wordle-in-terminal/helper"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func officialWordle(c *cli.Context) error {
	fmt.Println("official wordle")
	return nil
}
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
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		PaddingBottom(2).
		PaddingRight(2).
		PaddingTop(2).
		PaddingLeft(2).
		Width(22)
	fmt.Println(style.Render("hello world"))
	return
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
