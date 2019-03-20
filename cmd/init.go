package cmd

import (
	"os"

	"github.com/rrreeeyyy/giploc/container"

	"github.com/urfave/cli"
)

func InitCommand() cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "initialize the namespaces and launch the process (do not call it outside of giploc)",
		Action: func(context *cli.Context) error {
			if err := container.Initialization(); err != nil {
				os.Exit(1)
			}
			panic("container init failed to exec")
		},
	}
}
