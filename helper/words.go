package helper

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var Words = []string{"Quick", "Brown", "hello", "arose", "vivid"}

func Check(a string, b string) string {
	var str = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderLeft(true).
		BorderForeground(lipgloss.Color("#b673fa")).Render("h")
	fmt.Print(str)
	return str
}
func CheckAccuracy(solution string, guessed string) string {
	solutionWordArr := strings.Split(solution, "")
	guessedWordArr := strings.Split(guessed, "")
	finalWord := ""
	style := lipgloss.NewStyle().
		Bold(true).
		BorderLeft(true).
		BorderRight(true).
		BorderForeground(lipgloss.Color("63")).
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(2).
		PaddingRight(2)
	for i, value := range solutionWordArr {
		if guessedWordArr[i] == value {
			finalWord += style.
				Foreground(lipgloss.Color("#14c700")).
				Render(value)
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			finalWord += style.
				Foreground(lipgloss.Color("#e6d600")).
				Render(guessedWordArr[i])
			finalWord += lipgloss.PlaceHorizontal(10, lipgloss.Center, " ")
		}
		if !include(guessedWordArr[i], solutionWordArr) {
			finalWord += style.Render(guessedWordArr[i])
		}
	}
	return lipgloss.NewStyle().
		Margin(2).
		MarginLeft(5).
		Render(finalWord)
}

func include(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
