package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const byteCountInTestFile = 342190
const wordCountInTestFile = 58164
const lineCountInTestFile = 7145
const characterCountInTestFile = 339292

const TestFileName = "../data/wc_tool_test.txt"

func TestEwcWithOptionCFromFile(t *testing.T) {
	// given
	os.Args = []string{"ewc", "-c", TestFileName}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %s", byteCountInTestFile, TestFileName)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionLFromFile(t *testing.T) {
	// given
	os.Args = []string{"ewc", "-l", TestFileName}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %s", lineCountInTestFile, TestFileName)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionWFromFile(t *testing.T) {
	// given
	os.Args = []string{"ewc", "-w", TestFileName}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %s", wordCountInTestFile, TestFileName)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionMFromFile(t *testing.T) {
	// given
	os.Args = []string{"ewc", "-m", TestFileName}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %s", characterCountInTestFile, TestFileName)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithNoOptionFromFile(t *testing.T) {
	// given
	os.Args = []string{"ewc", TestFileName}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %d %d %s", lineCountInTestFile, wordCountInTestFile, byteCountInTestFile, TestFileName)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionCFromStdin(t *testing.T) {
	// given
	testInput := readTestDataFromWcToolTestFile()
	oldStdin := os.Stdin
	writeTestInputToStdin(testInput)
	defer func() { os.Stdin = oldStdin }()
	os.Args = []string{"ewc", "-c"}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d", 342190)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionLFromStdin(t *testing.T) {
	// given
	testInput := readTestDataFromWcToolTestFile()
	oldStdin := os.Stdin
	writeTestInputToStdin(testInput)
	defer func() { os.Stdin = oldStdin }()
	os.Args = []string{"ewc", "-l"}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d", 7145)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionWFromStdin(t *testing.T) {
	// given
	testInput := readTestDataFromWcToolTestFile()
	oldStdin := os.Stdin
	writeTestInputToStdin(testInput)
	defer func() { os.Stdin = oldStdin }()
	os.Args = []string{"ewc", "-w"}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d", 58164)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithOptionMFromStdin(t *testing.T) {
	// given
	testInput := readTestDataFromWcToolTestFile()
	oldStdin := os.Stdin
	writeTestInputToStdin(testInput)
	defer func() { os.Stdin = oldStdin }()
	os.Args = []string{"ewc", "-m"}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d", 339292)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func TestEwcWithNoOptionFromStdin(t *testing.T) {
	// given
	testInput := readTestDataFromWcToolTestFile()
	oldStdin := os.Stdin
	writeTestInputToStdin(testInput)
	defer func() { os.Stdin = oldStdin }()
	os.Args = []string{"ewc"}

	// when
	actualOutput := runMainAndCaptureOutput()

	// then
	expectedOutput := fmt.Sprintf("%d %d %d", 7145, 58164, 342190)
	if !strings.Contains(actualOutput, expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", actualOutput, expectedOutput)
	}
}

func readTestDataFromWcToolTestFile() string {
	testFile, _ := os.Open("../data/wc_tool_test.txt")
	testDataBytes, _ := io.ReadAll(testFile)
	testData := string(testDataBytes)
	return testData
}

func writeTestInputToStdin(testInput string) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdin = r
	go func() {
		_, err := w.Write([]byte(testInput))
		if err != nil {
			panic(err)
		}
		_ = w.Close()
	}()
}

func runMainAndCaptureOutput() string {
	// Keep backup of the real stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Close the Pipe so we can read it
	_ = w.Close()
	// Reset os.Stdout to its original value
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	return buf.String()
}
