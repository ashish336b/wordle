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
	case "grey":
		name = "#808080"
	case "cyan":
		name = "#42C2FF"
	}
	return baseStyle().BorderForeground(lipgloss.Color(name)).Foreground(lipgloss.Color(name))
}

func GetAlphabetKeyboard(name []string) string {
	finalString := ""
	for i, val := range name {
		i = i + 1
		finalString += lipgloss.NewStyle().PaddingLeft(2).BorderLeft(true).
			Bold(true).
			// BorderRight(true).
			// BorderForeground(lipgloss.Color("#e32289")).
			// BorderStyle(lipgloss.NormalBorder()).PaddingRight(5).
			Render(val)
		if i%7 == 0 {
			finalString += "\n"
		}

	}
	return lipgloss.NewStyle().MarginLeft(5).PaddingBottom(1).Render(finalString)
}

func PaintGreen(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#14c700")).Render(s)
}

func PaintYellow(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#e6d600")).Render(s)
}

func PaintGrey(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Render(s)
}

func instructionStyle() lipgloss.Style {
	return lipgloss.NewStyle().Margin(1).Bold(true).MarginLeft(3).Foreground(lipgloss.Color("#57f1ff"))
}

// TODO: optimize this function
func example(s string, b string, c bool) string {
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
	fmt.Println(example("AROSE", "A", true) + "\t" + "\"A\" is in the word and in the correct spot.\n")
	fmt.Println(example("PILLS", "I", false) + "\t" + "\"I\" is in the word and in the wrong spot.\n")
}
