package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nsf/termbox-go"
)


type Editor struct {
    width           int
    height          int
    editorXoffset   int
}

func InitEditor() *Editor {
    editor := Editor{}
    width, height := termbox.Size()
    editor.width = width * 90 / 100
    editor.height = height * 70 / 100
    editor.editorXoffset = ( width - editor.width )
    return &editor
}

func (m *Editor) Init() tea.Cmd {
    return nil
}

func (m *Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
            case "ctrl+c":
                return m, tea.Quit
        }
   }

    return m, nil
}

func (m *Editor) View() string {
    editorStyle := lipgloss.NewStyle().
                        Width(m.width * 50 /100).
                        Height(m.height * 65 / 100).
                        Padding(1).
                        Border(lipgloss.NormalBorder()).Align(lipgloss.Center)
    centeredEditor := lipgloss.NewStyle().PaddingLeft(m.editorXoffset).Render(editorStyle.Render(RenderMainPage()))
    return centeredEditor
}
