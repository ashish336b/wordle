package helper

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var Words = []string{"Quick", "Brown", "hello", "arose", "vivid"}

func Check() string {
	a := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#14c700")).
		Render("a")
	b := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#14c700")).
		Render("b")
	fmt.Println(a + b)
	return "hello"
}
func CheckAccuracy(solution string, guessed string) string {
	solutionWordArr := strings.Split(solution, "")
	guessedWordArr := strings.Split(guessed, "")
	finalWord := ""
	style := lipgloss.NewStyle().
		Bold(true).PaddingLeft(1)
	for i, value := range solutionWordArr {
		if guessedWordArr[i] == value {
			finalWord += style.
				Foreground(lipgloss.Color("#14c700")).
				Render(value)
			// finalWord += lipgloss.PlaceHorizontal(10, lipgloss.Center, " ")
		}
		if guessedWordArr[i] != value && include(guessedWordArr[i], solutionWordArr) {
			finalWord += style.
				Foreground(lipgloss.Color("#e6d600")).
				Render(guessedWordArr[i])
			finalWord += lipgloss.PlaceHorizontal(10, lipgloss.Center, " ")
		}
		if !include(guessedWordArr[i], solutionWordArr) {
			finalWord += style.Render(guessedWordArr[i])
			// finalWord += lipgloss.PlaceHorizontal(10, lipgloss.Center, " ")
		}
	}
	return lipgloss.
		NewStyle().
		Padding(1).
		Border(lipgloss.Border{
			Top:         "._.:*:",
			Bottom:      "._.:*:",
			Left:        "|*",
			Right:       "|*",
			TopLeft:     "*",
			TopRight:    "*",
			BottomLeft:  "*",
			BottomRight: "*",
		}).
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
