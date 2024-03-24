package main

import (
	"fmt"
	"os/exec"
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
	cmd := exec.Command("ewc", "-c", "wc", testFilename)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d %s", byteCountInTestFile, testFilename)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
	}
}

func TestWcWithOptionLWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	cmd := exec.Command("ewc", "-l", "wc", testFilename)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d %s", lineCountInTestFile, testFilename)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
	}
}

func TestWcWithOptionWWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	cmd := exec.Command("ewc", "-w", "wc", testFilename)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d %s", wordCountInTestFile, testFilename)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
	}
}

func TestWcWithOptionMWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	cmd := exec.Command("ewc", "-m", "wc", testFilename)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d %s", characterCountInTestFile, testFilename)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
	}
}

func TestWcWithNoOptionWithFile(t *testing.T) {
	// given
	testFilename := "wc_tool_test.txt"
	cmd := exec.Command("ewc", "wc", testFilename)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d %d %d %s", lineCountInTestFile, wordCountInTestFile, byteCountInTestFile, testFilename)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
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
//	if !strings.Contains(string(output), expectedOutput) {
//		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
//	}
//}

func TestWcWithOptionLWithoutFile(t *testing.T) {
	// given
	testFileLocation := "../../data/wc_tool_test.txt"
	commandString := fmt.Sprintf("cat %s | ewc -l wc", testFileLocation)
	cmd := exec.Command("sh", "-c", commandString)

	// when
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Command execution failed with error: %v", err)
	}

	// then
	expectedOutput := fmt.Sprintf("%d", lineCountInTestFile)
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Unexpected output. Got: %s, Expected: %s", string(output), expectedOutput)
	}
}
