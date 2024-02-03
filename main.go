package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main(){
    // check for command line arguments
    cmdArgs := os.Args
    hasFileName := false
    filePath := ""
    if(len(cmdArgs) == 2){
        hasFileName = true
        filePath = cmdArgs[1]
    }
    program := tea.NewProgram(InitEditor(hasFileName, filePath))
    if  _, err := program.Run(); err != nil {
        fmt.Println("Could not create a Bubble tea program!")
        os.Exit(-1)
    }
}
