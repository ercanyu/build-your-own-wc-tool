package handler

import (
	"fmt"
	"github.com/ercanyu/coding-challenges/wc-tool/internal/calculation"
)

type WcAction struct {
	Option   string
	Filename string
}

type WcActionType int

const (
	WcWithOptionC WcActionType = iota
	WcWithOptionL
	WcWithOptionW
	WcWithOptionM
	WcWithoutOption
)

func HandleWcAction(wcAction WcAction) (string, error) {
	filename := wcAction.Filename
	wcActionType := getWcActionType(wcAction)
	actionResult := ""

	switch wcActionType {
	case WcWithOptionC:
		actionResult = handleWcWithOptionC(filename)
	case WcWithOptionL:
		actionResult = handleWcWithOptionL(filename)
	case WcWithOptionW:
		actionResult = handleWcWithOptionW(filename)
	case WcWithOptionM:
		actionResult = handleWcWithOptionM(filename)
	case WcWithoutOption:
		actionResult = handleWcWithoutOption(filename)
	}
	return actionResult, nil
}

func handleWcWithoutOption(filename string) string {
	if filename != "" {
		numberOfBytes := calculation.WcCalculationFromFile(filename, calculation.NumberOfBytes)
		numberOfLines := calculation.WcCalculationFromFile(filename, calculation.NumberOfLines)
		numberOfWords := calculation.WcCalculationFromFile(filename, calculation.NumberOfWords)
		return fmt.Sprintf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, filename)
	} else {
		numberOfBytes := calculation.WcCalculationFromStdin(calculation.NumberOfBytes)
		numberOfLines := calculation.WcCalculationFromStdin(calculation.NumberOfLines)
		numberOfWords := calculation.WcCalculationFromStdin(calculation.NumberOfWords)
		return fmt.Sprintf("%d %d %d\n", numberOfLines, numberOfWords, numberOfBytes)
	}
}

func handleWcWithOptionM(filename string) string {
	var numberOfCharacters int
	if filename != "" {
		numberOfCharacters = calculation.WcCalculationFromFile(filename, calculation.NumberOfCharacters)
		return fmt.Sprintf("%d %s\n", numberOfCharacters, filename)
	} else {
		numberOfCharacters = calculation.WcCalculationFromStdin(calculation.NumberOfCharacters)
		return fmt.Sprintf("%d\n", numberOfCharacters)
	}
}

func handleWcWithOptionW(filename string) string {
	var numberOfWords int
	if filename != "" {
		numberOfWords = calculation.WcCalculationFromFile(filename, calculation.NumberOfWords)
		return fmt.Sprintf("%d %s\n", numberOfWords, filename)
	} else {
		numberOfWords = calculation.WcCalculationFromStdin(calculation.NumberOfWords)
		return fmt.Sprintf("%d\n", numberOfWords)
	}
}

func handleWcWithOptionL(filename string) string {
	var numberOfLines int
	if filename != "" {
		numberOfLines = calculation.WcCalculationFromFile(filename, calculation.NumberOfLines)
		return fmt.Sprintf("%d %s\n", numberOfLines, filename)
	} else {
		numberOfLines = calculation.WcCalculationFromStdin(calculation.NumberOfLines)
		return fmt.Sprintf("%d\n", numberOfLines)
	}
}

func handleWcWithOptionC(filename string) string {
	var numberOfBytes int
	if filename != "" {
		numberOfBytes = calculation.WcCalculationFromFile(filename, calculation.NumberOfBytes)
		return fmt.Sprintf("%d %s\n", numberOfBytes, filename)
	} else {
		numberOfBytes = calculation.WcCalculationFromStdin(calculation.NumberOfBytes)
		return fmt.Sprintf("%d\n", numberOfBytes)
	}
}

func getWcActionType(wcAction WcAction) WcActionType {
	switch wcAction.Option {
	case "c":
		return WcWithOptionC
	case "l":
		return WcWithOptionL
	case "w":
		return WcWithOptionW
	case "m":
		return WcWithOptionM
	default:
		return WcWithoutOption
	}
}
