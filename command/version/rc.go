package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
)

func rcCmd() cli.Command {
	return cli.Command {
		Name: "rc",
		Usage: "Process a release candidate version - set RC version, build, test, commit, tag and push",
		ArgsUsage: "VERSION RC_NUMBER",
		Action: func(ctx *cli.Context) error {
			return rcRun(ctx)
		},
	}
}

func rcRun(ctx *cli.Context) error {
	version := ctx.Args().Get(0)
	number := ctx.Args().Get(1)
	rcVersion := version + "-rc" + number
	fmt.Printf("Release candidate version: %s", rcVersion)
	return nil
}