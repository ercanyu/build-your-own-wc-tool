package calculation

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

const DataDirectory = "/data/"

type WcCalculationType int

const (
	NumberOfLines WcCalculationType = iota
	NumberOfBytes
	NumberOfWords
	NumberOfCharacters
)

func WcCalculation(filename string, option WcCalculationType) int {
	switch option {
	case NumberOfLines:
		return calculateNumberOfLines(filename)
	case NumberOfBytes:
		return calculateNumberOfBytesIn(filename)
	case NumberOfWords:
		return calculateNumberOfWords(filename)
	case NumberOfCharacters:
		return calculateNumberOfCharactersInFile(filename)
	default:
		panic(fmt.Sprintf("Invalid option: %d", option))
	}
}

func calculateNumberOfCharactersInFile(filename string) int {
	file, err := openFile(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfCharacters := countNumberOfCharacters(file)
	closeFile(file)
	return numberOfCharacters
}

func countNumberOfCharacters(file *os.File) int {
	numberOfCharacters := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		numberOfCharacters++
	}
	return numberOfCharacters
}

func calculateNumberOfWords(filename string) int {
	file, err := openFile(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfWords := countNumberOfWords(file)
	closeFile(file)
	return numberOfWords
}

func countNumberOfWords(file *os.File) int {
	numberOfWords := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numberOfWords++
	}
	return numberOfWords
}

func calculateNumberOfLines(fileName string) int {
	var numberOfLines int
	if fileName == "" {
		num := 0
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			num++
		}
		return num
	} else {
		file, err := openFile(fileName)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return 0
		}
		numberOfLines = countNumberOfLines(file)
		closeFile(file)
	}
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

func calculateNumberOfBytesIn(fileName string) int {
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfBytes := countNumberOfBytes(file)
	closeFile(file)
	return numberOfBytes
}

func countNumberOfBytes(file *os.File) int {
	var sizeInBytes int
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file: ", err)
			}
			break
		}
		sizeInBytes += n
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
