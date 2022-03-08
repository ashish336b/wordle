package helper

import "github.com/charmbracelet/lipgloss"

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
