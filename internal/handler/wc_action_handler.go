package handler

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal/calculation"
	"io"
	"os"
)

type WcAction struct {
	OptionFlag string
	FileName   string
}

type WcActionType int
type wcActionHandler func(s string) string

const (
	WcWithOptionC WcActionType = iota
	WcWithOptionL
	WcWithOptionW
	WcWithOptionM
	WcWithoutOption
)

var wcActionHandlers = map[WcActionType]wcActionHandler{
	WcWithOptionC:   handleWcWithOptionC,
	WcWithOptionL:   handleWcWithOptionL,
	WcWithOptionW:   handleWcWithOptionW,
	WcWithOptionM:   handleWcWithOptionM,
	WcWithoutOption: handleWcWithoutOption,
}

func HandleWcAction(wcAction WcAction) (string, error) {
	filename := wcAction.FileName
	wcActionType := getWcActionType(wcAction.OptionFlag)
	actionResult := ""

	wcHandler, ok := wcActionHandlers[wcActionType]
	if !ok {
		return "", fmt.Errorf("invalid wcAction Option: %s", wcAction.OptionFlag)
	}
	actionResult = wcHandler(filename)

	return actionResult, nil
}

func handleWcWithoutOption(filename string) string {
	if filename != "" {
		numberOfBytes := calculation.WcCalculationFromFile(filename, calculation.NumberOfBytes)
		numberOfLines := calculation.WcCalculationFromFile(filename, calculation.NumberOfLines)
		numberOfWords := calculation.WcCalculationFromFile(filename, calculation.NumberOfWords)
		return fmt.Sprintf("%d %d %d", numberOfLines, numberOfWords, numberOfBytes)
	} else {
		input := createStringFromStdin()
		numberOfBytes := calculation.WcCalculationFromString(input, calculation.NumberOfBytes)
		numberOfLines := calculation.WcCalculationFromString(input, calculation.NumberOfLines)
		numberOfWords := calculation.WcCalculationFromString(input, calculation.NumberOfWords)
		return fmt.Sprintf("%d %d %d", numberOfLines, numberOfWords, numberOfBytes)
	}
}

func handleWcWithOptionM(filename string) string {
	var numberOfCharacters int
	if filename != "" {
		numberOfCharacters = calculation.WcCalculationFromFile(filename, calculation.NumberOfCharacters)
		return fmt.Sprintf("%d", numberOfCharacters)
	} else {
		input := createStringFromStdin()
		numberOfCharacters = calculation.WcCalculationFromString(input, calculation.NumberOfCharacters)
		return fmt.Sprintf("%d", numberOfCharacters)
	}
}

func handleWcWithOptionW(filename string) string {
	var numberOfWords int
	if filename != "" {
		numberOfWords = calculation.WcCalculationFromFile(filename, calculation.NumberOfWords)
		return fmt.Sprintf("%d", numberOfWords)
	} else {
		input := createStringFromStdin()
		numberOfWords = calculation.WcCalculationFromString(input, calculation.NumberOfWords)
		return fmt.Sprintf("%d", numberOfWords)
	}
}

func handleWcWithOptionL(filename string) string {
	var numberOfLines int
	if filename != "" {
		numberOfLines = calculation.WcCalculationFromFile(filename, calculation.NumberOfLines)
		return fmt.Sprintf("%d", numberOfLines)
	} else {
		input := createStringFromStdin()
		numberOfLines = calculation.WcCalculationFromString(input, calculation.NumberOfLines)
		return fmt.Sprintf("%d", numberOfLines)
	}
}

func handleWcWithOptionC(filename string) string {
	var numberOfBytes int
	if filename != "" {
		numberOfBytes = calculation.WcCalculationFromFile(filename, calculation.NumberOfBytes)
		return fmt.Sprintf("%d", numberOfBytes)
	} else {
		input := createStringFromStdin()
		numberOfBytes = calculation.WcCalculationFromString(input, calculation.NumberOfBytes)
		return fmt.Sprintf("%d", numberOfBytes)
	}
}

func createStringFromStdin() string {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)
	return input
}

func getWcActionType(optionFlag string) WcActionType {
	switch optionFlag {
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
