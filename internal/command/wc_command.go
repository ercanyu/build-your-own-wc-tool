package command

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal"
	"github.com/ercanyu/build-your-own-wc-tool/internal/handler"
	ufcli "github.com/urfave/cli/v2"
	"io"
	"os"
)

func RunWcCommand(ctx *ufcli.Context) error {
	fileName := ctx.Args().Get(0)
	var reader io.Reader
	if fileName != "" {
		file, err := internal.OpenFile(fileName)
		if err != nil {
			fmt.Printf("Failed to open file %s\n", fileName)
		}
		reader = file
	} else {
		reader = os.Stdin
	}
	wcAction := handler.WcAction{
		OptionFlag: getOptionFromContext(ctx),
		Reader:     reader,
	}
	actionResult, err := handler.HandleWcAction(wcAction)

	if fileName != "" {
		fmt.Printf("%s %s\n", actionResult, fileName)
	} else {
		fmt.Printf("%s\n", actionResult)
	}

	return err
}
func getOptionFromContext(context *ufcli.Context) string {
	if context.Bool("c") {
		return "c"
	}
	if context.Bool("l") {
		return "l"
	}
	if context.Bool("w") {
		return "w"
	}
	if context.Bool("m") {
		return "m"
	}
	return ""
}
