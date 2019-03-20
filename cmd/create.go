package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rrreeeyyy/giploc/container"
	"github.com/rrreeeyyy/giploc/oci"

	"github.com/urfave/cli"
)

// CreateCommand ...
func CreateCommand() cli.Command {
	return cli.Command{
		Name:      "create",
		Usage:     "create a container",
		ArgsUsage: "<container-id>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "bundle, b",
				Value: "",
				Usage: `path to the root of the bundle directory, defaults to the current directory`,
			},
			cli.StringFlag{
				Name:  "console-socket",
				Value: "",
				Usage: "path to an AF_UNIX socket which will receive a file descriptor referencing the master end of the console's pseudoterminal",
			},
			cli.StringFlag{
				Name:  "pid-file",
				Value: "",
				Usage: "specify the file to write the process id to",
			},
			cli.BoolFlag{
				Name:  "no-pivot",
				Usage: "do not use pivot root to jail process inside rootfs.  This should be used whenever the rootfs is on top of a ramdisk",
			},
			cli.BoolFlag{
				Name:  "no-new-keyring",
				Usage: "do not create a new session keyring for the container.  This will cause the container to inherit the calling processes session key",
			},
			cli.IntFlag{
				Name:  "preserve-fds",
				Usage: "Pass N additional file descriptors to the container (stdio + $LISTEN_FDS + N in total)",
			},
		},
		Action: func(context *cli.Context) error {
			fmt.Println("args: %v", context.Args())

			spec, err := oci.LoadSpec(filepath.Join(context.String("bundle"), "config.json"))
			if err != nil {
				return err
			}
			status, err := container.Create(context, spec)
			if err != nil {
				return err
			}

			os.Exit(status)
			return nil
		},
	}
}
