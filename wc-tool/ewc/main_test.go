package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

const byteCountInTestFile = 342190
const wordCountInTestFile = 58164
const lineCountInTestFile = 7145
const characterCountInTestFile = 339292

func TestWcWithOptionCWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	os.Args = []string{"ewc", "-c", "wc", testFilename}

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %s", byteCountInTestFile, testFilename)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestWcWithOptionLWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	os.Args = []string{"ewc", "-l", "wc", testFilename}

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %s", lineCountInTestFile, testFilename)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestWcWithOptionWWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	os.Args = []string{"ewc", "-w", "wc", testFilename}

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %s", wordCountInTestFile, testFilename)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestWcWithOptionMWithFileWithCapture(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	os.Args = []string{"ewc", "-m", "wc", testFilename}

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %s", characterCountInTestFile, testFilename)
	if !strings.Contains(string(actualOutput), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestWcWithNoOptionWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	os.Args = []string{"ewc", "wc", testFilename}

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %d %d %s", lineCountInTestFile, wordCountInTestFile, byteCountInTestFile, testFilename)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

//func TestWcWithOptionCWithoutFile(t *testing.T) {
//	// given
//	testFileLocation := "../../data/wc_tool_test.txt"
//	commandString := fmt.Sprintf("cat %s | ewc -c wc", testFileLocation)
//	cmd := exec.Command("sh", "-c", commandString)
//
//	// when
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		t.Errorf("Command execution failed with error: %v", err)
//	}
//
//	// then
//	expectedOutput := fmt.Sprintf("%d", byteCountInTestFile)
//	if !strings.Contains(actualOutput, expectedOutput) {
//		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
//	}
//}

//func TestWcWithOptionLWithoutFile(t *testing.T) {
//	// given
//	testFileLocation := "../../data/wc_tool_test.txt"
//	commandString := fmt.Sprintf("cat %s | ewc -l wc", testFileLocation)
//	cmd := exec.Command("sh", "-c", commandString)
//
//	// when
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		t.Errorf("Command execution failed with error: %v", err)
//	}
//
//	// then
//	expectedOutput := fmt.Sprintf("%d", lineCountInTestFile)
//	if !strings.Contains(actualOutput, expectedOutput) {
//		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
//	}
//}

func TestWcWithNoOptionWithStdin(t *testing.T) {
	// given
	testInput := "this is a test\nthis is not a test though\n"
	expectedLineCount := 2
	expectedWordCount := 10
	expectedByteCount := len(testInput)
	os.Args = []string{"ewc", "wc"}

	oldStdin := os.Stdin
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	go func() {
		_, err := w.Write([]byte(testInput))
		if err != nil {
			t.Error(err)
		}
		w.Close()
	}()

	// when
	actualOutput := runAndCaptureOutput(main)

	// then
	expectedOutput := fmt.Sprintf("%d %d %d", expectedLineCount, expectedWordCount, expectedByteCount)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func runAndCaptureOutput(f func()) string {
	// Keep backup of the real stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// Close the Pipe so we can read it
	_ = w.Close()
	// Reset os.Stdout to its original value
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	return buf.String()
}
