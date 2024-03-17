package handler

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const DataDirectory = "/data/"

func HandleWcCommand(filename string, option string) int {
	if option == "l" {
		return findNumberOfLinesInFile(filename)
	} else if option == "c" {
		return int(findNumberOfBytesInFile(filename))
	}
	return 0
}

func findNumberOfLinesInFile(fileName string) int {
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfLines := countNumberOfLines(file)
	defer closeFile(file)
	return numberOfLines
}

func countNumberOfLines(file *os.File) int {
	numberOfLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines
}

func findNumberOfBytesInFile(fileName string) int64 {
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfBytes := countNumberOfBytes(file)
	defer closeFile(file)
	return numberOfBytes
}

func countNumberOfBytes(file *os.File) int64 {
	var sizeInBytes int64
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file: ", err)
			}
			break
		}
		sizeInBytes += int64(n)
	}
	return sizeInBytes
}

func openFile(fileName string) (*os.File, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, err
	}
	fullFilePath := projectRoot + DataDirectory + fileName
	file, err := os.Open(fullFilePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func closeFile(file *os.File) {
	_ = file.Close()
}

func findProjectRoot() (string, error) {
	markerFileName := "go.mod"
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, markerFileName)); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("marker file '%s' not found in the directory tree", markerFileName)
		}

		currentDir = parentDir
	}
}
