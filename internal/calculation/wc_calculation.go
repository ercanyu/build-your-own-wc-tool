package calculation

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type WcCalculationType int

type wcCalculation func(reader io.Reader) []int

const NewLineLength = 2 //len([]byte("\r\n"))

const (
	NumberOfLines WcCalculationType = iota
	NumberOfBytes
	NumberOfWords
	NumberOfCharacters
	NumberOfLinesWordsBytes
)

var wcCalculators = map[WcCalculationType]wcCalculation{
	NumberOfLines:           findNumberOfLines,
	NumberOfBytes:           findNumberOfBytes,
	NumberOfWords:           findNumberOfWords,
	NumberOfCharacters:      findNumberOfCharacters,
	NumberOfLinesWordsBytes: findNumberOfLinesWordsBytes,
}

func WcCalculation(reader io.Reader, wcCalculationType WcCalculationType) []int {
	wcCalculator, ok := wcCalculators[wcCalculationType]
	if !ok {
		panic(fmt.Sprintf("Invalid wcCalculationType: %d", wcCalculationType))
	}

	return wcCalculator(reader)
}

func findNumberOfBytes(reader io.Reader) []int {
	numberOfBytes := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		numberOfBytes += len(bytes) + NewLineLength
	}

	return []int{numberOfBytes}
}

func findNumberOfLines(reader io.Reader) []int {
	numberOfLines := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		numberOfLines++
	}
	return []int{numberOfLines}
}

func findNumberOfWords(reader io.Reader) []int {
	numberOfWords := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		numberOfWords++
	}
	return []int{numberOfWords}
}

func findNumberOfCharacters(reader io.Reader) []int {
	numberOfCharacters := 0
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		numberOfCharacters++
	}
	return []int{numberOfCharacters}
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
