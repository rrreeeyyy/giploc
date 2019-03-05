package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	exitCodeSuccess = iota
	exitCodeError
	exitCodeMisuse
)

const (
	appName   = "giploc"
	appConfig = "config.json"
	appUsage  = `Minimal container runtime`
)

func Run(args []string) int {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage

	app.Commands = []cli.Command{
		StartCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeError)
	}

	return 0
}
