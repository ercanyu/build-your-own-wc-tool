package handler

import "testing"

func TestHandleWcActionWithOptionC(t *testing.T) {
	// given
	wcAction := WcAction{
		Option:   "c",
		Filename: "wc_tool_test.txt",
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
		Option:   "l",
		Filename: "wc_tool_test.txt",
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
		Option:   "w",
		Filename: "wc_tool_test.txt",
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
		Option:   "m",
		Filename: "wc_tool_test.txt",
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
