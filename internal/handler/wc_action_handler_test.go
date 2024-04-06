package handler

import (
	"github.com/ercanyu/build-your-own-wc-tool/internal"
	"os"
	"testing"
)

func TestHandleWcActionWithOptionC(t *testing.T) {
	// given
	reader := prepareReader(t)
	wcAction := WcAction{
		OptionFlag: "c",
		Reader:     reader,
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "342190"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
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

func TestHandleWcActionWithOptionL(t *testing.T) {
	// given
	reader := prepareReader(t)
	wcAction := WcAction{
		OptionFlag: "l",
		Reader:     reader,
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionW(t *testing.T) {
	// given
	reader := prepareReader(t)
	wcAction := WcAction{
		OptionFlag: "w",
		Reader:     reader,
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "58164"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionM(t *testing.T) {
	// given
	reader := prepareReader(t)
	wcAction := WcAction{
		OptionFlag: "m",
		Reader:     reader,
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "339292"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithoutOption(t *testing.T) {
	// given
	reader := prepareReader(t)
	wcAction := WcAction{
		Reader: reader,
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 58164 342190"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}
