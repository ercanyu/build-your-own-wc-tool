package command

import (
	"fmt"
	"github.com/ercanyu/wc-tool/internal/calculation"
	ufcli "github.com/urfave/cli/v2"
)

func WcCommand() *ufcli.Command {
	return &ufcli.Command{
		Name:  "wc",
		Usage: "wc tool command",
		Action: func(ctx *ufcli.Context) error {
			return handleWcAction(ctx)
		},
	}
}

func handleWcAction(ctx *ufcli.Context) error {
	filename := ctx.Args().Get(0)
	if ctx.Bool("c") {
		numberOfBytes := calculation.HandleWcCalculation(filename, calculation.NumberOfBytes)
		fmt.Printf("%d %s\n", numberOfBytes, filename)
	} else if ctx.Bool("l") {
		numberOfLines := calculation.HandleWcCalculation(filename, calculation.NumberOfLines)
		fmt.Printf("%d %s\n", numberOfLines, filename)
	} else if ctx.Bool("w") {
		numberOfWords := calculation.HandleWcCalculation(filename, calculation.NumberOfWords)
		fmt.Printf("%d %s\n", numberOfWords, filename)
	} else if ctx.Bool("m") {
		numberOfCharacters := calculation.HandleWcCalculation(filename, calculation.NumberOfCharacters)
		fmt.Printf("%d %s\n", numberOfCharacters, filename)
	} else {
		numberOfBytes := calculation.HandleWcCalculation(filename, calculation.NumberOfBytes)
		numberOfLines := calculation.HandleWcCalculation(filename, calculation.NumberOfLines)
		numberOfWords := calculation.HandleWcCalculation(filename, calculation.NumberOfWords)
		fmt.Printf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, filename)
	}
	return nil
}
