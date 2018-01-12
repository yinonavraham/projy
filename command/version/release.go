package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
)

func releaseCmd() cli.Command {
	return cli.Command {
		Name: "release",
		Usage: "Process a release version - set version, build, test, commit, tag and push",
		ArgsUsage: "VERSION",
		Action: func(ctx *cli.Context) error {
			return releaseRun(ctx)
		},
	}
}

func releaseRun(ctx *cli.Context) error {
	version := ctx.Args().Get(0)
	fmt.Printf("Release version: %s", version)
	return nil
}