package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Window struct {
	screens []Screen
}

func InitWindow() *Window {
	window := Window{
		screens: []Screen{NewWelcomeScreen()},
	}

	return &window
}

func (e *Editor) UpdateWindow(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e *Editor) ViewWindow() string {
	return e.window.screens[e.currentActiveScreen].render()
}

type ScreenType int

const (
	WELCOME_SCREEN ScreenType = iota
	HELP_SCREEN
	EDITOR_SCREEN
	BUFFER_LIST_SCREEN
)

type Screen interface {
	render() string
}

type WelcomeScreen struct {
	appName         string
	instructions    string
	backgroundColor string
	foregroundColor string
}

func NewWelcomeScreen() *WelcomeScreen {

	titleStr := "  _____ _   _ _______  _______ \n" +
		" |_   _| | | | ____\\ \\/ / ____|\n" +
		"   | | | |_| |  _|  \\  /|  _|  \n" +
		"   | | |  _  | |___ /  \\| |___ \n" +
		"   |_| |_| |_|_____/_/\\_\\_____|\n\n" +
		"******* A Terminal Hex Editor ********\n"
	instructionsStr := "\n" +
        GetNormalText("Type ") + GetStyledCode("[") + GetNormalText(" anytime to toggle the editor to COMMAND mode.") +"\n\n" +
        GetNormalText("Type ") + GetStyledCode(":h<Enter>") + GetNormalText(" to open the manual.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":q<Enter>") + GetNormalText(" to quit the application when in COMMAND mode.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":e <File_Path><Enter>") + GetNormalText(" in COMMAND mode to open and start editing a file.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":w<Enter>") + GetNormalText(" inside open file to save the changes.") + "\n\n"

	return &WelcomeScreen{
		appName:         titleStr,
		instructions:    instructionsStr,
		backgroundColor: "103",
		foregroundColor: "234",
	}
}

func (s *WelcomeScreen) render() string {
	screenStyle := lipgloss.NewStyle().
		Width(100).
		Height(30).
		Border(lipgloss.NormalBorder()).
		Background(lipgloss.Color(s.backgroundColor)).
		Foreground(lipgloss.Color(s.foregroundColor)).
		Align(lipgloss.Center).PaddingTop(5)

	instructionsStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(s.backgroundColor)).
		Foreground(lipgloss.Color(s.foregroundColor)).
        Align(lipgloss.Center).
		MarginTop(2)

	return screenStyle.Render(s.appName + "\n" + instructionsStyle.Render(s.instructions))
}
