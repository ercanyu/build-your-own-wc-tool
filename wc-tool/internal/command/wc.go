package command

import (
	"fmt"
	"github.com/ercanyu/wc-tool/internal/handler"
	ufcli "github.com/urfave/cli/v2"
)

func WcCommand() *ufcli.Command {
	return &ufcli.Command{
		Name:  "wc",
		Usage: "wc tool",
		Action: func(ctx *ufcli.Context) error {
			return handleWcAction(ctx)
		},
	}
}

func handleWcAction(ctx *ufcli.Context) error {
	filename := ctx.Args().Get(0)
	if ctx.Bool("c") {
		numberOfBytes := handler.HandleWcCommand(filename, "c")
		fmt.Printf("%d %s\n", numberOfBytes, filename)
	} else if ctx.Bool("l") {
		numberOfLines := handler.HandleWcCommand(filename, "l")
		fmt.Printf("%d %s\n", numberOfLines, filename)
	}
	return nil
}
