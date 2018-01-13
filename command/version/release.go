package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/yinonavraham/projy/flag"
)

func releaseCmd() cli.Command {
	return cli.Command {
		Name: "release",
		Usage: "Process a release version - set version, build, test, commit, tag and push",
		ArgsUsage: "VERSION",
		Flags: []cli.Flag {
			flag.ForceFlag,
			flag.DryRunFlag,
			flag.SkipTestsFlag,
			flag.SkipTagFlag,
			flag.SkipPushFlag,
		},
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