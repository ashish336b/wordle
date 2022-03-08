package helper

import (
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
