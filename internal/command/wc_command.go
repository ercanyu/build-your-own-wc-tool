package command

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal/handler"
	ufcli "github.com/urfave/cli/v2"
)

func RunWcCommand(ctx *ufcli.Context) error {
	wcAction := handler.WcAction{
		OptionFlag: getOptionFromContext(ctx),
		Filename:   ctx.Args().Get(0),
	}
	actionResult, err := handler.HandleWcAction(wcAction)
	fmt.Print(actionResult)
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
