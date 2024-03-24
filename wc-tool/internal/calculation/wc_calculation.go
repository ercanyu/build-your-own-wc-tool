package calculation

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const DataDirectory = "/data/"

type WcCalculationType int

const (
	NumberOfLines WcCalculationType = iota
	NumberOfBytes
	NumberOfWords
	NumberOfCharacters
)

func WcCalculationFromFile(filename string, option WcCalculationType) int {
	switch option {
	case NumberOfBytes:
		return findNumberOfBytesInFile(filename)
	case NumberOfLines:
		return findNumberOfLinesInFile(filename)
	case NumberOfWords:
		return findNumberOfWordsInFile(filename)
	case NumberOfCharacters:
		return findNumberOfCharactersInFile(filename)
	default:
		panic(fmt.Sprintf("Invalid option: %d", option))
	}
}

func WcCalculationFromStdin(option WcCalculationType) int {
	switch option {
	case NumberOfBytes:
		return findNumberOfBytesFromStdin()
	case NumberOfLines:
		return findNumberOfLinesFromStdin()
	case NumberOfWords:
		return findNumberOfWordsFromStdin()
	case NumberOfCharacters:
		return findNumberOfCharactersFromStdin()
	default:
		panic(fmt.Sprintf("Invalid option: %d", option))
	}
}

func WcCalculationFromString(input string, option WcCalculationType) int {
	switch option {
	case NumberOfBytes:
		return findNumberOfBytesFromString(input)
	case NumberOfLines:
		return findNumberOfLinesFromString(input)
	case NumberOfWords:
		return findNumberOfWordsFromString(input)
	case NumberOfCharacters:
		return findNumberOfCharactersFromString(input)
	default:
		panic(fmt.Sprintf("Invalid option: %d", option))
	}
}

func findNumberOfBytesFromString(input string) int {
	return len([]byte(input))
}

func findNumberOfLinesFromString(input string) int {
	return strings.Count(input, "\n")
}

func findNumberOfWordsFromString(input string) int {
	return len(strings.Fields(input))
}

func findNumberOfCharactersFromString(input string) int {
	return len([]rune(input))
}

func findNumberOfBytesInFile(fileName string) int {
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	scanner := bufio.NewScanner(file)
	numberOfBytes := countNumberOfBytesFromScanner(scanner)
	closeFile(file)
	return numberOfBytes
}

func findNumberOfLinesInFile(fileName string) int {
	var numberOfLines int
	file, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	scanner := bufio.NewScanner(file)
	numberOfLines = countNumberOfLinesFromScanner(scanner)
	closeFile(file)
	return numberOfLines
}

func findNumberOfWordsInFile(filename string) int {
	file, err := openFile(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	scanner := bufio.NewScanner(file)
	numberOfWords := countNumberOfWordsFromScanner(scanner)
	closeFile(file)
	return numberOfWords
}

func findNumberOfCharactersInFile(filename string) int {
	file, err := openFile(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	scanner := bufio.NewScanner(file)
	numberOfCharacters := countNumberOfCharactersFromScanner(scanner)
	closeFile(file)
	return numberOfCharacters
}

func findNumberOfBytesFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	return countNumberOfBytesFromScanner(scanner)
}

func findNumberOfLinesFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	return countNumberOfLinesFromScanner(scanner)
}

func findNumberOfWordsFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	return countNumberOfWordsFromScanner(scanner)
}

func findNumberOfCharactersFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	return countNumberOfCharactersFromScanner(scanner)
}

func countNumberOfBytesFromScanner(scanner *bufio.Scanner) int {
	numberOfBytes := 0
	newLineLength := len([]byte("\r\n"))
	for scanner.Scan() {
		bytes := scanner.Bytes()
		numberOfBytes += len(bytes) + newLineLength
	}

	return numberOfBytes
}

func countNumberOfLinesFromScanner(scanner *bufio.Scanner) int {
	numberOfLines := 0
	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines
}

func countNumberOfWordsFromScanner(scanner *bufio.Scanner) int {
	numberOfWords := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numberOfWords++
	}
	return numberOfWords
}

func countNumberOfCharactersFromScanner(scanner *bufio.Scanner) int {
	numberOfCharacters := 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		numberOfCharacters++
	}
	return numberOfCharacters
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
