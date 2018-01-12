package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
)

func currentCmd() cli.Command {
	return cli.Command {
		Name: "current",
		Usage: "Print the current project version",
		Action: func(ctx *cli.Context) error {
			return currentRun(ctx)
		},
	}
}

func currentRun(ctx *cli.Context) error {
	fmt.Printf("Current project version: ???")
	return nil
}