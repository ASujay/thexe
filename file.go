package main

import (
	"fmt"
	"os"
)

type File struct {
	filePath   string
	fileHandle *os.File
	bytes      []byte
	curPos     int
}

func InitFile(filePath string) File {
	// open the file
	fileHandle, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		os.Exit(1)
	}

	defer fileHandle.Close()

	// reading the content of the file
	fileInfo, err := fileHandle.Stat()
	if err != nil {
		fmt.Println("Error getting the file info: ", err)
		os.Exit(1)
	}

	// get the filesize
	fileSize := fileInfo.Size()
	fileBytes := make([]byte, fileSize)

	_, err = fileHandle.Read(fileBytes)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	file := File{}
	file.filePath = filePath
	file.bytes = fileBytes
	file.curPos = 0
	file.fileHandle = fileHandle
	return file
}

func (f File) Close() {
	f.fileHandle.Close()
}

func (f File) Debug() {
	// Print the bytes in hexadecimal format
	for i, byteValue := range f.bytes {
		if i%16 == 0 {
			fmt.Printf("\n%08X: %02X ", i, byteValue)
		} else {
			fmt.Printf("%02X ", byteValue)
		}
	}
	fmt.Println()
}
