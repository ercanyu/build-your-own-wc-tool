package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const ToolName = "wc"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter commands(type 'exit' to quit): ")

	for {
		fmt.Print("wc-tool> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}
		command = command[:len(command)-1]

		if command == "exit" {
			fmt.Println("Exiting...")
			break
		}

		commandParts := strings.Split(command, " ")
		if len(commandParts) < 3 || commandParts[0] != ToolName {
			printInvalidCommandMessage()
			continue
		}
		commandToExecute := commandParts[1]

		if commandToExecute == "-c" {
			fileName := commandParts[2]
			sizeInBytes := findFileSizeInBytes(fileName)
			fmt.Printf("%d %s\n", sizeInBytes, fileName)
		} else {
			printInvalidCommandMessage()
		}
	}
}

func findFileSizeInBytes(fileName string) int64 {
	file := openFile(fileName)
	sizeInBytes := calculateSizeInBytes(file)
	closeFile(file)
	return sizeInBytes
}

func calculateSizeInBytes(file *os.File) int64 {
	var sizeInBytes int64
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("Error reading file: ", err)
			break
		}
		sizeInBytes += int64(n)
	}
	return sizeInBytes
}

func openFile(fileName string) *os.File {
	fullFilePath := getParentDirectory() + "/data/" + fileName
	file, err := os.Open(fullFilePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Error closing file: ", err)
	}
}

func getParentDirectory() string {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory: ", err)
		return ""
	}

	return filepath.Join(currentWorkingDirectory, "..")
}

func printInvalidCommandMessage() {
	fmt.Printf("Invalid command, exc. '%s -c <file-name>'\n", ToolName)
}
