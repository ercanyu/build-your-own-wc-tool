package handler

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal/calculation"
	"io"
)

type WcAction struct {
	OptionFlag string
	Reader     io.Reader
}

type WcActionType int
type wcActionHandler func(reader io.Reader) string

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
	wcActionType := getWcActionType(wcAction.OptionFlag)
	actionResult := ""

	wcHandler, ok := wcActionHandlers[wcActionType]
	if !ok {
		return "", fmt.Errorf("invalid wcAction Option: %s", wcAction.OptionFlag)
	}
	actionResult = wcHandler(wcAction.Reader)

	return actionResult, nil
}

func handleWcWithoutOption(reader io.Reader) string {
	numberOfLinesWordsBytes := calculation.WcCalculation(reader, calculation.NumberOfLinesWordsBytes)
	return fmt.Sprintf(
		"%d %d %d",
		numberOfLinesWordsBytes[0],
		numberOfLinesWordsBytes[1],
		numberOfLinesWordsBytes[2],
	)
}

func handleWcWithOptionM(reader io.Reader) string {
	numberOfCharacters := calculation.WcCalculation(reader, calculation.NumberOfCharacters)
	return fmt.Sprintf("%d", numberOfCharacters[0])
}

func handleWcWithOptionW(reader io.Reader) string {
	numberOfWords := calculation.WcCalculation(reader, calculation.NumberOfWords)
	return fmt.Sprintf("%d", numberOfWords[0])
}

func handleWcWithOptionL(reader io.Reader) string {
	numberOfLines := calculation.WcCalculation(reader, calculation.NumberOfLines)
	return fmt.Sprintf("%d", numberOfLines[0])
}

func handleWcWithOptionC(reader io.Reader) string {
	numberOfBytes := calculation.WcCalculation(reader, calculation.NumberOfBytes)
	return fmt.Sprintf("%d", numberOfBytes[0])
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
