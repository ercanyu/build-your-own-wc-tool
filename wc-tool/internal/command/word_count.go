package command

import (
	"bufio"
	"fmt"
	ufcli "github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

const DataDirectory = "/data/"

func WordCountCommand() *ufcli.Command {
	return &ufcli.Command{
		Name:  "wc",
		Usage: "wc tool",
		Action: func(ctx *ufcli.Context) error {
			filename := ctx.Args().Get(0)
			if ctx.Bool("c") {
				numberOfBytes := findNumberOfBytesInFile(filename)
				fmt.Printf("%d %s\n", numberOfBytes, filename)
			} else if ctx.Bool("l") {
				numberOfLines := findNumberOfLinesInFile(filename)
				fmt.Printf("%d %s\n", numberOfLines, filename)
			}

			return nil
		},
	}
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
