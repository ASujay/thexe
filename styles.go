package main

import "github.com/charmbracelet/lipgloss"

func GetCommandLineElementStyle(width int, backgroundColor string, foregroundColor string) lipgloss.Style {
	cmdElementStyle := lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.NormalBorder()).
		Background(lipgloss.Color(backgroundColor)).
		Foreground(lipgloss.Color(foregroundColor))
	return cmdElementStyle
}

func GetStyledCode(text string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("246")).
		Foreground(lipgloss.Color("53")).Render(text)
}

func GetNormalText(text string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("103")).
		Foreground(lipgloss.Color("234")).
		Render(text)
}
