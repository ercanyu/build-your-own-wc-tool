package calculation

import (
	"bufio"
	"io"
	"strings"
)

type WcCalculationResult struct {
	ByteCount      int
	CharacterCount int
	WordCount      int
	LineCount      int
}

const NewLineByteCount = 2      // len([]byte("\r\n"))
const NewLineCharacterCount = 2 // \r and \n

func WcCalculation(reader io.Reader) WcCalculationResult {
	scanner := bufio.NewScanner(reader)
	var lineCount, wordCount, byteCount, characterCount int
	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		wordCount += len(strings.Fields(line))
		byteCount += len(line) + NewLineByteCount
		characterCount += len([]rune(line)) + NewLineCharacterCount
	}

	return WcCalculationResult{
		ByteCount:      byteCount,
		CharacterCount: characterCount,
		WordCount:      wordCount,
		LineCount:      lineCount,
	}
}
