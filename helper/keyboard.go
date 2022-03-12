package helper

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var Alphabets = []string{"A", "B", "C", "D", "E", "F", "G", "H",
	"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

var GreenWord = []string{}
var YellowWord = []string{}
var GrayWord = []string{}

func getAlphabetIndex(alphabet string) int {
	alphabets := []string{"A", "B", "C", "D", "E", "F", "G", "H",
		"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return index(strings.ToUpper(alphabet), alphabets)
}
func GetAlphabetKeyboard(name []string) string {
	for _, val := range GreenWord {
		Alphabets[getAlphabetIndex(val)] = PaintGreen(val)
	}
	for _, val := range YellowWord {
		Alphabets[getAlphabetIndex(val)] = PaintYellow(val)
	}
	for _, val := range GrayWord {
		Alphabets[getAlphabetIndex(val)] = PaintGrey(val)
	}
	finalString := ""
	for i, val := range name {
		i = i + 1
		finalString += lipgloss.NewStyle().PaddingLeft(2).BorderLeft(true).
			Bold(true).
			Render(val)
		if i%7 == 0 {
			finalString += "\n"
		}

	}
	return lipgloss.NewStyle().MarginLeft(5).PaddingBottom(1).Render(finalString)
}
