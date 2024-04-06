package calculation

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type WcCalculationType int

const (
	NumberOfLines WcCalculationType = iota
	NumberOfBytes
	NumberOfWords
	NumberOfCharacters
	NumberOfLinesWordsBytes
)

const NewLineLength = 2 //len([]byte("\r\n"))

func WcCalculation(reader io.Reader, option WcCalculationType) []int {
	switch option {
	case NumberOfBytes:
		return []int{findNumberOfBytes(reader)}
	case NumberOfLines:
		return []int{findNumberOfLines(reader)}
	case NumberOfWords:
		return []int{findNumberOfWords(reader)}
	case NumberOfCharacters:
		return []int{findNumberOfCharacters(reader)}
	case NumberOfLinesWordsBytes:
		return findNumberOfLinesWordsBytes(reader)
	default:
		panic(fmt.Sprintf("Invalid option: %d", option))
	}
}

func findNumberOfBytes(reader io.Reader) int {
	numberOfBytes := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		numberOfBytes += len(bytes) + NewLineLength
	}

	return numberOfBytes
}

func findNumberOfLines(reader io.Reader) int {
	numberOfLines := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		numberOfLines++
	}
	return numberOfLines
}

func findNumberOfWords(reader io.Reader) int {
	numberOfWords := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numberOfWords++
	}
	return numberOfWords
}

func findNumberOfCharacters(reader io.Reader) int {
	numberOfCharacters := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		numberOfCharacters++
	}
	return numberOfCharacters
}

func findNumberOfLinesWordsBytes(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	var lines, words, bytes int
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		bytes += len(line) + NewLineLength
	}

	return []int{lines, words, bytes}
}
