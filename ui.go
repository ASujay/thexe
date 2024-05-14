package main

import "github.com/charmbracelet/lipgloss"


func RenderCode(code string) string {
    return lipgloss.NewStyle().Foreground(lipgloss.Color(80)).Render(code)
}

func RenderMainPage() string {
    title := " _____ _                   \n" + 
             "|_   _| |__   _____  _____ \n" +
             "   | | | '_ \\ / _ \\ \\/ / _ \\ \n" +
             "  | | | | | |  __/>  <  __/\n" +
             "    |_| |_| |_|\\___/_/\\_\\___\\ \n"
    option := "\n\n\n\n\nEnter " + RenderCode(":w<Enter>") + " to save the current file.\n\n\n" +
              "Enter " + RenderCode(":q<Enter>") + " to exit the editor.\n\n\n" +
              "Enter " + RenderCode(":b<Enter>") + " to open the buffer list.\n\n\n" +
              "Enter " + RenderCode(":e <filename><Enter>") + " to edit a file.\n\n\n"
    return title + option
}
