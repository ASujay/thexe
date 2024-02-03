package main

import (
	"github.com/charmbracelet/lipgloss"
)

type ScreenType int

const (
    WELCOME ScreenType = iota
    MANUAL
    BUFFERLIST
    HEXEDITOR
)

type Screen interface{
    render() string
}

type WelcomeScreen struct {
    backgroundColor string
    textColor string
    screenBuffer string
    width int
    height int
}

func getWelcomeScreenBuffer(backgroundColor string, textColor string) string  {
	titleStr := "  _____ _   _ _______  _______ \n" +
		" |_   _| | | | ____\\ \\/ / ____|\n" +
		"   | | | |_| |  _|  \\  /|  _|  \n" +
		"   | | |  _  | |___ /  \\| |___ \n" +
		"   |_| |_| |_|_____/_/\\_\\_____|\n\n" +
		"******* A Terminal Hex Editor ********\n"
	instructionsStr := "\n" +
		GetNormalText("Type ") + GetStyledCode("[") + GetNormalText(" anytime to toggle the editor to COMMAND mode.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":h<Enter>") + GetNormalText(" to open the manual.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":q<Enter>") + GetNormalText(" to quit the application when in COMMAND mode.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":e <File_Path><Enter>") + GetNormalText(" in COMMAND mode to open and start editing a file.") + "\n\n" +
		GetNormalText("Type ") + GetStyledCode(":w<Enter>") + GetNormalText(" inside open file to save the changes.") + "\n\n"

    screenStyle := lipgloss.NewStyle().
		Width(100).
		Height(30).
		Border(lipgloss.NormalBorder()).
		Background(lipgloss.Color(backgroundColor)).
		Foreground(lipgloss.Color(textColor)).
		Align(lipgloss.Center).PaddingTop(5)

	instructionsStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(backgroundColor)).
		Foreground(lipgloss.Color(textColor)).
		Align(lipgloss.Center).
		MarginTop(2)

	return screenStyle.Render(titleStr + "\n" +  instructionsStyle.Render(instructionsStr))
	
}

func InitWelcomeScreen() WelcomeScreen {
    welcomeScreen := WelcomeScreen{
        backgroundColor: "103",
        textColor: "234",
        width: 100,
        height: 30,
    }

    welcomeScreen.screenBuffer = getWelcomeScreenBuffer(
        welcomeScreen.backgroundColor, welcomeScreen.textColor,
    )
    return welcomeScreen
}

func (s WelcomeScreen) render() string {
    return s.screenBuffer
}


// Hex Editor Screen

type HexScreen struct {
    filepath string
    textColor string
    screenBuffer string
    width int
    height int

}

func (s HexScreen) render() string {
   return "Hex Screen" 
}

func InitHexScreen(filepath string) HexScreen {
    return HexScreen{
        filepath: filepath,
    }
}
