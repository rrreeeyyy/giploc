package main

import (
	"os"

	"github.com/rrreeeyyy/giploc/cmd"
)

func main() {
	exit := cmd.Run(os.Args)
	os.Exit(exit)
}
