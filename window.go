package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)


/****************************************************************************
 *                                   Tab                                    * 
 ****************************************************************************
*/

// This is window types
const (
	WelcomeScreen = iota
	HexEditor
)

// the editor can have multiple tabs open
type Tab struct {
	File       File
	WindowType int
}


func (tab Tab) Render() string {
    return "View"
}

/****************************************************************************
 *                               Command Line                               * 
 ****************************************************************************
*/

type CommandLine struct {
    model tea.Model 
}

func (commandLine CommandLine) Render() string {
    return "Command Line"
}

/****************************************************************************
 *                                  Editor                                  * 
 ****************************************************************************
*/

func InitEditor(args []string) Editor {
	editor := Editor{
        activeTab: 0,
        commandLine: CommandLine{},
        tabs: []Tab{
            Tab{},
        },
    }
	return editor
}

type Editor struct {
    tabs []Tab
    activeTab uint
    commandLine CommandLine
}

func (editor Editor) Init() tea.Cmd {
    return nil
}

func (editor Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
        case tea.KeyMsg:
            switch msg.String() {
                case "ctrl+c", "q":
                    return editor, tea.Quit
            }
    }

    return editor, nil
}

func (editor Editor) View() string {
    // we have 2 diferent things that we render.
    // first is the actual window which will be either the main page
    // or some binary file
    // second would be the command line / status line which will be at
    // the bottom of the window
    // get the string representation of the tab 
    activeTab := editor.tabs[editor.activeTab].Render()
    // get the string representation of the command line
    commandLine := editor.commandLine.Render()
    return activeTab + "\n" + commandLine + "\n" 
}

func Run(editor *Editor) {
    program := tea.NewProgram(*editor)
    if _, err := program.Run(); err != nil {
        fmt.Printf("Program cannot start, Error: %v\n", err)
        os.Exit(1)
    }    
}


