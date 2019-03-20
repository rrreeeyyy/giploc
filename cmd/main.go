package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	version string
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

// Run ...
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = version
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output for logging",
		},
		cli.StringFlag{
			Name:  "log",
			Value: "/dev/null",
			Usage: "set the log file path where internal debug information is written",
		},
		cli.StringFlag{
			Name:  "log-format",
			Value: "text",
			Usage: "set the format used by logs ('text' (default), or 'json')",
		},
		cli.StringFlag{
			Name:  "root",
			Value: "/run/giploc",
			Usage: "root directory for storage of container state (this should be located in tmpfs)",
		},
	}

	app.Commands = []cli.Command{
		InitCommand(),
		StateCommand(),
		StartCommand(),
		CreateCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeError)
	}

	return 0
}
