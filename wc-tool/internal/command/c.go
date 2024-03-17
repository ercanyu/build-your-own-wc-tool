package command

import (
	"fmt"
	ufcli "github.com/urfave/cli/v2"
)

func NewCCommand() *ufcli.Command {
	return &ufcli.Command{
		Name:  "c",
		Usage: "size in bytes",
		Action: func(c *ufcli.Context) error {
			fmt.Println("c command")
			return nil
		},
	}
}
