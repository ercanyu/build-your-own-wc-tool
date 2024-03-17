package command

import (
	"fmt"
	"github.com/ercanyu/wc-tool/internal/handler"
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
		numberOfBytes := handler.HandleWcCalculation(filename, handler.NumberOfBytes)
		fmt.Printf("%d %s\n", numberOfBytes, filename)
	} else if ctx.Bool("l") {
		numberOfLines := handler.HandleWcCalculation(filename, handler.NumberOfLines)
		fmt.Printf("%d %s\n", numberOfLines, filename)
	} else if ctx.Bool("w") {
		numberOfWords := handler.HandleWcCalculation(filename, handler.NumberOfWords)
		fmt.Printf("%d %s\n", numberOfWords, filename)
	} else if ctx.Bool("m") {
		numberOfCharacters := handler.HandleWcCalculation(filename, handler.NumberOfCharacters)
		fmt.Printf("%d %s\n", numberOfCharacters, filename)
	} else {
		numberOfBytes := handler.HandleWcCalculation(filename, handler.NumberOfBytes)
		numberOfLines := handler.HandleWcCalculation(filename, handler.NumberOfLines)
		numberOfWords := handler.HandleWcCalculation(filename, handler.NumberOfWords)
		fmt.Printf("%d %d %d %s\n", numberOfLines, numberOfWords, numberOfBytes, filename)
	}
	return nil
}
