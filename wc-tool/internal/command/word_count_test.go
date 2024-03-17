package command

import (
	"testing"
)

func TestFindNumberOfLinesInFile(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := findNumberOfLinesInFile(filename)

	// then
	if numberOfLines != 7145 {
		t.Errorf("Expected 7145, got %d", numberOfLines)
	}
}

func TestFindNumberOfBytesInFile(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfBytes := findNumberOfBytesInFile(filename)

	// then
	if numberOfBytes != 342190 {
		t.Errorf("Expected 342190, got %d", numberOfBytes)
	}
}
