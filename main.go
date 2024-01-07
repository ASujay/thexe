package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(InitEditor())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
}
