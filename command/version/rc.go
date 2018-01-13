package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/yinonavraham/projy/flag"
)

func rcCmd() cli.Command {
	return cli.Command {
		Name: "rc",
		Usage: "Process a release candidate version - set RC version, build, test, commit, tag and push",
		ArgsUsage: "VERSION RC_NUMBER",
		Flags: []cli.Flag{
			flag.ForceFlag,
			flag.DryRunFlag,
			flag.SkipTestsFlag,
			flag.SkipTagFlag,
			flag.SkipPushFlag,
		},
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