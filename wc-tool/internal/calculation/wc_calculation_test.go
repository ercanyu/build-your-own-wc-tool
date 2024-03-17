package calculation

import (
	"testing"
)

func TestWcCalculationForNumberOfBytes(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfBytes := WcCalculation(filename, NumberOfBytes)

	// then
	if numberOfBytes != 342190 {
		t.Errorf("Expected 342190, got %d", numberOfBytes)
	}
}

func TestWcCalculationForNumberOfLines(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := WcCalculation(filename, NumberOfLines)

	// then
	if numberOfLines != 7145 {
		t.Errorf("Expected 7145, got %d", numberOfLines)
	}
}

func TestWcCalculationForNumberOfWords(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := WcCalculation(filename, NumberOfWords)

	// then
	if numberOfLines != 58164 {
		t.Errorf("Expected 58164, got %d", numberOfLines)
	}
}

func TestWcCalculationForNumberOfCharacters(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := WcCalculation(filename, NumberOfCharacters)

	// then
	if numberOfLines != 339292 {
		t.Errorf("Expected 339292, got %d", numberOfLines)
	}
}
