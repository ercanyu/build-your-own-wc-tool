package handler

import "testing"

func TestHandleWcActionWithOptionC(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "c",
		FileName:   "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "342190"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionL(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "l",
		FileName:   "wc_tool_test.txt",
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
	wcAction := WcAction{
		OptionFlag: "w",
		FileName:   "wc_tool_test.txt",
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
	wcAction := WcAction{
		OptionFlag: "m",
		FileName:   "wc_tool_test.txt",
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
	wcAction := WcAction{
		FileName: "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 58164 342190"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}
