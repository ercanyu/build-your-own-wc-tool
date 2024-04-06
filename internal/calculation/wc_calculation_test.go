package calculation

import (
	"github.com/ercanyu/build-your-own-wc-tool/internal"
	"os"
	"testing"
)

func TestWcCalculationForNumberOfBytes(t *testing.T) {
	// given
	reader := prepareReader(t)

	// when
	numberOfBytes := WcCalculation(reader, NumberOfBytes)

	// then
	if numberOfBytes[0] != 342190 {
		t.Errorf("Expected 342190, got %d", numberOfBytes)
	}
}

func TestWcCalculationForNumberOfLines(t *testing.T) {
	// given
	reader := prepareReader(t)

	// when
	numberOfLines := WcCalculation(reader, NumberOfLines)

	// then
	if numberOfLines[0] != 7145 {
		t.Errorf("Expected 7145, got %d", numberOfLines)
	}
}

func TestWcCalculationForNumberOfWords(t *testing.T) {
	// given
	reader := prepareReader(t)

	// when
	numberOfLines := WcCalculation(reader, NumberOfWords)

	// then
	if numberOfLines[0] != 58164 {
		t.Errorf("Expected 58164, got %d", numberOfLines)
	}
}

func TestWcCalculationForNumberOfCharacters(t *testing.T) {
	// given
	reader := prepareReader(t)

	// when
	numberOfLines := WcCalculation(reader, NumberOfCharacters)

	// then
	if numberOfLines[0] != 339292 {
		t.Errorf("Expected 339292, got %d", numberOfLines)
	}
}

func prepareReader(t *testing.T) *os.File {
	filename := "wc_tool_test.txt"
	reader, err := internal.OpenFile(filename)
	if err != nil {
		t.Errorf("could not open file %s", filename)
	}
	return reader
}
