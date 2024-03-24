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
	WcCommandWithOptionC WcActionType = iota
	WcCommandWithOptionL
	WcCommandWithOptionW
	WcCommandWithOptionM
	WcCommandWithoutOption
)

func HandleWcAction(wcAction WcAction) (string, error) {
	filename := wcAction.Filename
	wcActionType := getWcActionType(wcAction)
	actionResult := ""

	switch wcActionType {
	case WcCommandWithOptionC:
		numberOfBytes := calculation.WcCalculation(filename, calculation.NumberOfBytes)
		actionResult = fmt.Sprintf("%d %s\n", numberOfBytes, filename)
	case WcCommandWithOptionL:
		numberOfLines := calculation.WcCalculation(filename, calculation.NumberOfLines)
		actionResult = fmt.Sprintf("%d %s\n", numberOfLines, filename)
	case WcCommandWithOptionW:
		numberOfWords := calculation.WcCalculation(filename, calculation.NumberOfWords)
		actionResult = fmt.Sprintf("%d %s\n", numberOfWords, filename)
	case WcCommandWithOptionM:
		numberOfCharacters := calculation.WcCalculation(filename, calculation.NumberOfCharacters)
		actionResult = fmt.Sprintf("%d %s\n", numberOfCharacters, filename)
	case WcCommandWithoutOption:
		numberOfBytes := calculation.WcCalculation(filename, calculation.NumberOfBytes)
		numberOfLines := calculation.WcCalculation(filename, calculation.NumberOfLines)
		numberOfWords := calculation.WcCalculation(filename, calculation.NumberOfWords)
		actionResult = fmt.Sprintf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, filename)
	}
	return actionResult, nil
}

func getWcActionType(wcAction WcAction) WcActionType {
	switch wcAction.Option {
	case "c":
		return WcCommandWithOptionC
	case "l":
		return WcCommandWithOptionL
	case "w":
		return WcCommandWithOptionW
	case "m":
		return WcCommandWithOptionM
	default:
		return WcCommandWithoutOption
	}
}
