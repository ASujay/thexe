package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Buffer struct {

}

type Mode int

const (
    COMMAND Mode = iota
    NORMAL
    EDIT
)


type UIElement int

const (
    WINDOW UIElement = iota
    COMMANDLINE
    FILESTATUS
)

func InitEditor (hasFileName bool, filePath string) Editor {
    editor := Editor{}
    editor.currentMode = COMMAND
    editor.activeComponent = COMMANDLINE
    editor.currentScreen = WELCOME
    editor.bufferMap = make(map[string]Screen)
    editor.bufferMap["welcome"] = InitWelcomeScreen()
    editor.currentScreenTitle = "welcome"
    if hasFileName {
       // then we open the file and create the hex editor screen
       editor.bufferMap[filePath] = InitHexScreen(filePath) 
       editor.currentScreen = HEXEDITOR
       editor.currentScreenTitle = filePath
    }
    return editor
}

type Editor struct {
    bufferMap map[string]Screen
    currentMode Mode
    activeComponent UIElement
    currentScreen ScreenType
    currentScreenTitle string
}


func (e Editor) Init() tea.Cmd {
    return nil
}

func (e Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
        case tea.KeyMsg: {
            switch msg.String() {
                case "ctrl+c", "q": 
                    return e, tea.Quit
            }
        }
    }

    return e, nil
}

func (e Editor) View() string {
    return e.bufferMap[e.currentScreenTitle].render() + "\n"
}

