package main

import "os"

func main() {
	editor := InitEditor(os.Args)
	Run(&editor)
}
