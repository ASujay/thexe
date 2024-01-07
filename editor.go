package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Editor struct {
	width               int
	height              int
	cmdLine             *CmdLine
	status              Status
	mode                Mode
	window              *Window
	currentActiveScreen ScreenType
}

type Mode int

const (
	NORMAL Mode = iota
	COMMAND
)

var (
	TERMINAL_WIDTH  int
	TERMINAL_HEIGHT int
)

func (m Mode) String() string {
	switch m {
	case NORMAL:
		return "NORMAL"
	case COMMAND:
		return "COMMAND"
	default:
		return "UNKNOWN MODE"
	}
}

func InitEditor() *Editor {
	initialMode := NORMAL
	editor := &Editor{
		width:               60,
		height:              40,
		cmdLine:             InitCmdLine(),
		mode:                initialMode,
		currentActiveScreen: WELCOME_SCREEN,
		window:              InitWindow(),
		status: Status{
			row:      "X",
			col:      "X",
			fileName: "",
			modeStr:  initialMode.String(),
		},
	}
	return editor
}

func (e *Editor) Init() tea.Cmd {
	return nil
}

func (e *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		TERMINAL_WIDTH = msg.Width
		TERMINAL_HEIGHT = msg.Height
		return e, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return e, tea.Quit
		case "[":
			e.mode = COMMAND
			return e, nil
		default:
			if e.mode == COMMAND {
				return e.UpdateCmdLine(msg)
			}
			return e, nil
		}
	}
	return e, nil
}

func (e *Editor) View() string {
	return e.ViewWindow() + "\n" + e.ViewCmdLine() + "\n"
}
