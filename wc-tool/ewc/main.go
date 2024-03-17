package main

import (
	"github.com/ercanyu/wc-tool/internal/command"
	ufcli "github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cliApp := ufcli.App{
		Name:    "ewc",
		Version: "1.0.0",
		Usage:   "A simple word count tool",
		Flags: []ufcli.Flag{
			&ufcli.BoolFlag{
				Name:    "bytes",
				Aliases: []string{"c"},
				Usage:   "size in bytes",
			},
			&ufcli.BoolFlag{
				Name:    "lines",
				Aliases: []string{"l"},
				Usage:   "count lines",
			},
			&ufcli.BoolFlag{
				Name:    "words",
				Aliases: []string{"w"},
				Usage:   "count words",
			},
			&ufcli.BoolFlag{
				Name:    "characters",
				Aliases: []string{"m"},
				Usage:   "count characters",
			},
		},
		Commands: []*ufcli.Command{
			command.WcCommand(),
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
