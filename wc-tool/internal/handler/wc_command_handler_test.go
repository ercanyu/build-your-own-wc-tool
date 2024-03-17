package handler

import (
	"testing"
)

func TestHandleWcCommandWithOptionC(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfBytes := HandleWcCommand(filename, "c")

	// then
	if numberOfBytes != 342190 {
		t.Errorf("Expected 342190, got %d", numberOfBytes)
	}
}

func TestHandleWcCommandWithOptionL(t *testing.T) {
	// given
	filename := "test.txt"

	// when
	numberOfLines := HandleWcCommand(filename, "l")

	// then
	if numberOfLines != 7145 {
		t.Errorf("Expected 7145, got %d", numberOfLines)
	}
}
