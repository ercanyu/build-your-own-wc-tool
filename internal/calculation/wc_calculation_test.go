package calculation

import (
	"os"
	"testing"
)

func TestWcCalculation(t *testing.T) {
	// given
	filename := "../../data/wc_tool_test.txt"
	reader, err := os.Open(filename)
	if err != nil {
		t.Errorf("could not open file %s", filename)
	}

	// when
	wcCalculationResult := WcCalculation(reader)

	// then
	expectedWcCalculationResult := WcCalculationResult{
		ByteCount:      342190,
		LineCount:      7145,
		WordCount:      58164,
		CharacterCount: 339292,
	}
	if wcCalculationResult != expectedWcCalculationResult {
		t.Errorf(
			"Expected %v, got %d",
			expectedWcCalculationResult,
			wcCalculationResult)
	}
}
