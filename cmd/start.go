package cmd

import (
	"os"

	"github.com/rrreeeyyy/giploc/container"
	"github.com/urfave/cli"
)

func StartCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "executes the user defined process",
		Action: func(context *cli.Context) error {
			spec, err := container.SetupSpec(context)
			if err != nil {
				return err
			}
			code, err := container.Start(context, spec)
			if err != nil {
				return err
			}
			os.Exit(code)
			return nil
		},
	}
}
