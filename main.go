package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nsf/termbox-go"
)

func main() {
    err := termbox.Init()
    if err != nil {
        fmt.Println("Failed to initialize termbox: ", err)
        os.Exit(1)
    }
    program:= tea.NewProgram(InitEditor())
    if _, err := program.Run(); err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }
    termbox.Close()
}
