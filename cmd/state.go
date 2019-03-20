package cmd

import (
	"encoding/json"
	"os"
	"time"

	"github.com/urfave/cli"
)

type containerState struct {
	Version        string            `json:"ociVersion"`
	ID             string            `json:"id"`
	InitProcessPid int               `json:"pid"`
	Status         string            `json:"status"`
	Bundle         string            `json:"bundle"`
	Rootfs         string            `json:"rootfs"`
	Created        time.Time         `json:"created"`
	Annotations    map[string]string `json:"annotations,omitempty"`
	Owner          string            `json:"owner"`
}

func StateCommand() cli.Command {
	return cli.Command{
		Name:  "state",
		Usage: "state of a container",
		Action: func(context *cli.Context) error {
			cs := containerState{
				Version:        "1.0.1",
				ID:             context.Args().Get(0),
				InitProcessPid: 23174,
				Status:         "created",
				Rootfs:         "/home/ubuntu/alpine-bundle/rootfs",
				Created:        time.Now(),
			}
			data, err := json.MarshalIndent(cs, "", "  ")
			if err != nil {
				return err
			}
			os.Stdout.Write(data)
			return nil
		},
	}
}
