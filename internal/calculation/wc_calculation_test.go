package calculation

import (
	"github.com/ercanyu/build-your-own-wc-tool/internal"
	"testing"
)

func TestWcCalculation(t *testing.T) {
	// given
	filename := "wc_tool_test.txt"
	reader, err := internal.OpenFile(filename)
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
