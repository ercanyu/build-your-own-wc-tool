package handler

import "testing"

func TestHandleWcActionWithOptionC(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "c",
		Filename:   "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "342190 wc_tool_test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionL(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "l",
		Filename:   "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 wc_tool_test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionW(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "w",
		Filename:   "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "58164 wc_tool_test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionM(t *testing.T) {
	// given
	wcAction := WcAction{
		OptionFlag: "m",
		Filename:   "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "339292 wc_tool_test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithoutOption(t *testing.T) {
	// given
	wcAction := WcAction{
		Filename: "wc_tool_test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 58164 342190 wc_tool_test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}
