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
	if option == NumberOfLines {
		return findNumberOfLinesInFile(filename)
	} else if option == NumberOfBytes {
		return findNumberOfBytesInFile(filename)
	} else if option == NumberOfWords {
		return findNumberOfWordsInFile(filename)
	} else if option == NumberOfCharacters {
		return findNumberOfCharactersInFile(filename)
	}
	panic(fmt.Sprintf("Invalid option: %d", option))
}

func findNumberOfCharactersInFile(filename string) int {
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

func findNumberOfWordsInFile(filename string) int {
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

func findNumberOfLinesInFile(fileName string) int {
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	numberOfLines := countNumberOfLines(file)
	closeFile(file)
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

func findNumberOfBytesInFile(fileName string) int {
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
