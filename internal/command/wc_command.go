package command

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal"
	"github.com/ercanyu/build-your-own-wc-tool/internal/calculation"
	ufcli "github.com/urfave/cli/v2"
	"io"
	"os"
)

const OptionFlagForBytes = "c"
const OptionFlagForLines = "l"
const OptionFlagForWords = "w"
const OptionFlagForCharacters = "m"

func RunWcCommand(ctx *ufcli.Context) error {
	fileName := ctx.Args().Get(0)
	reader, err := createReader(fileName)
	if err != nil {
		return err
	}

	wcCalculationResult := calculation.WcCalculation(reader)
	resultMessage := getResultMessage(ctx, wcCalculationResult)
	if fileName != "" {
		fmt.Printf("%s %s\n", resultMessage, fileName)
	} else {
		fmt.Printf("%s\n", resultMessage)
	}

	return nil
}

func createReader(fileName string) (io.Reader, error) {
	var reader io.Reader
	if fileName != "" {
		file, err := internal.OpenFile(fileName)
		if err != nil {
			fmt.Printf("Failed to open file %s\n", fileName)
			return nil, err
		}
		reader = file
	} else {
		reader = os.Stdin
	}
	return reader, nil
}

func getResultMessage(
	context *ufcli.Context,
	wcCalculationResult calculation.WcCalculationResult,
) string {
	if context.Bool(OptionFlagForBytes) {
		return fmt.Sprintf("%d", wcCalculationResult.ByteCount)
	}
	if context.Bool(OptionFlagForLines) {
		return fmt.Sprintf("%d", wcCalculationResult.LineCount)
	}
	if context.Bool(OptionFlagForWords) {
		return fmt.Sprintf("%d", wcCalculationResult.WordCount)
	}
	if context.Bool(OptionFlagForCharacters) {
		return fmt.Sprintf("%d", wcCalculationResult.CharacterCount)
	}

	return fmt.Sprintf(
		"%d %d %d",
		wcCalculationResult.LineCount,
		wcCalculationResult.WordCount,
		wcCalculationResult.ByteCount,
	)
}
