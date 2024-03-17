package handler

import "testing"

func TestHandleWcActionWithOptionC(t *testing.T) {
	// given
	wcAction := WcAction{
		Option:   "c",
		Filename: "test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "342190 test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionL(t *testing.T) {
	// given
	wcAction := WcAction{
		Option:   "l",
		Filename: "test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionW(t *testing.T) {
	// given
	wcAction := WcAction{
		Option:   "w",
		Filename: "test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "58164 test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithOptionM(t *testing.T) {
	// given
	wcAction := WcAction{
		Option:   "m",
		Filename: "test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "339292 test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}

func TestHandleWcActionWithoutOption(t *testing.T) {
	// given
	wcAction := WcAction{
		Filename: "test.txt",
	}

	// when
	actionResult, _ := HandleWcAction(wcAction)

	// then
	expectedResult := "7145 58164 342190 test.txt\n"
	if actionResult != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, actionResult)
	}
}
