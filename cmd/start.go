package cmd

import (
	"os"
	"path/filepath"

	"github.com/rrreeeyyy/giploc/container"
	"github.com/rrreeeyyy/giploc/oci"

	"github.com/urfave/cli"
)

// StartCommand ...
func StartCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "executes the user defined process",
		Action: func(context *cli.Context) error {
			spec, err := oci.LoadSpec(filepath.Join(context.String("bundle"), "config.json"))
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
