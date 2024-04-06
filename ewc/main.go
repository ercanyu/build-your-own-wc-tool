package main

import (
	"fmt"
	"github.com/ercanyu/build-your-own-wc-tool/internal/command"
	ufcli "github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	cliApp := ufcli.App{
		Name:    "ewc",
		Version: "1.0.0",
		Usage:   "ewc -c|l|w|m [filename]",
		Flags: []ufcli.Flag{
			&ufcli.BoolFlag{
				Name:    "bytes",
				Aliases: []string{command.OptionFlagForBytes},
				Usage:   "size in bytes",
			},
			&ufcli.BoolFlag{
				Name:    "lines",
				Aliases: []string{command.OptionFlagForLines},
				Usage:   "count lines",
			},
			&ufcli.BoolFlag{
				Name:    "words",
				Aliases: []string{command.OptionFlagForWords},
				Usage:   "count words",
			},
			&ufcli.BoolFlag{
				Name:    "characters",
				Aliases: []string{command.OptionFlagForCharacters},
				Usage:   "count characters",
			},
		},
		Action: func(ctx *ufcli.Context) error {
			return command.RunWcCommand(ctx)
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Now().Sub(start).Microseconds()
	fmt.Printf("completed in %.3f ms\n", float32(elapsed)/1000.0)
}
