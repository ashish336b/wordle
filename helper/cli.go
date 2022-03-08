package helper

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func baseStyle() lipgloss.Style {
	style := lipgloss.NewStyle().
		BorderLeft(true).
		Bold(true).
		BorderRight(true).
		BorderForeground(lipgloss.Color("63")).
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(2).
		PaddingRight(2)
	return style
}

func coloredString(name string) lipgloss.Style {
	switch strings.ToLower(name) {
	case "green":
		name = "#14c700"
	case "yellow":
		name = "#e6d600"
	case "cyan":
		name = "#42C2FF"
	}
	return baseStyle().Foreground(lipgloss.Color(name))
}

func instructionStyle() lipgloss.Style {
	return lipgloss.NewStyle().Margin(1).MarginLeft(3).Foreground(lipgloss.Color("#57f1ff"))
}

// TODO: optimize this function
func Example(s string, b string, c bool) string {
	ex1 := strings.Split(strings.ToUpper(s), "")
	str1 := ""
	for _, val := range ex1 {
		style := lipgloss.NewStyle().MarginLeft(3)
		if val == strings.ToUpper(b) {
			if c {
				str1 += style.Foreground(lipgloss.Color("#14c700")).Render(val)
			} else {
				str1 += style.Foreground(lipgloss.Color("#e6d600")).Render(val)
			}
		} else {
			str1 += style.Render(val)
		}
	}
	return str1
}

func ShowInstruction() {
	fmt.Println(instructionStyle().Render("Instruction!"))
	text := instructionStyle().
		UnsetMarginTop().
		Render("Guess the WORDLE in six tries.\nEach guess must be a valid five-letter word.\nHit the enter button to submit\nAfter each guess, the color of the word will change to show how close your guess was to the word.")
	fmt.Println(text)
	fmt.Println(instructionStyle().
		UnsetMarginTop().
		Foreground(lipgloss.Color("#fff")).Render("Example!"))
	fmt.Println(Example("AROSE", "A", true) + "\t" + "\"A\" is in the word and in the correct spot.\n")
	fmt.Println(Example("PILLS", "I", false) + "\t" + "\"A\" is in the word and in the wrong spot.\n")
}
