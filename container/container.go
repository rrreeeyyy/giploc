package container

import (
	"github.com/urfave/cli"
)

type Container interface {
	ID() string
	Start(process *Process) (err error)
	Run(process *Process) (err error)
}

func Start(cli *cli.Context, args string) (int, error) {
	return 0, nil
}

func SetupSpec(context *cli.Context) (string, error) {
	return "", nil
}
