package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Status struct {
	modeStr  string
	row      string
	col      string
	fileName string
}

type CmdLine struct {
	textInput textinput.Model
}

func (e *Editor) UpdateCmdLine(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "enter":
				e.cmdLine.textInput.SetValue("")
				e.mode = NORMAL
				return e, nil
			}
		}
	}
	e.cmdLine.textInput, cmd = e.cmdLine.textInput.Update(msg)
	return e, cmd
}

func (e *Editor) ViewCmdLine() string {
	cmdStyle := GetCommandLineElementStyle(
		55, "109", "232",
	)

	modeStyle := GetCommandLineElementStyle(
		11, "109", "232",
	).Align(lipgloss.Center)

	colAndRowStyle := GetCommandLineElementStyle(
		30, "109", "232",
	).Align(lipgloss.Center)

	return lipgloss.JoinHorizontal(
		lipgloss.Right,
		cmdStyle.Render(e.cmdLine.textInput.View()),
		modeStyle.Render(e.mode.String()),
		colAndRowStyle.Render("( "+e.status.row, ", "+e.status.col+" )"),
	)
}

func InitCmdLine() *CmdLine {
	input := textinput.New()
	input.Focus()
	input.CharLimit = 50
	input.Width = 100
	input.Cursor.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("254"))
	return &CmdLine{
		textInput: input,
	}
}
