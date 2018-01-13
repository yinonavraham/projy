package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/yinonavraham/projy/flag"
)

func snapshotCmd() cli.Command {
	return cli.Command {
		Name: "snapshot",
		Usage: "Process a snapshot version - set version, build, test, commit and push",
		ArgsUsage: "VERSION",
		Flags: []cli.Flag{
			flag.ForceFlag,
			flag.DryRunFlag,
			flag.SkipTestsFlag,
			flag.SkipTagFlag,
			flag.SkipPushFlag,
			QualifierFlag,
		},
		Action: func(ctx *cli.Context) error {
			return snapshotRun(ctx)
		},
	}
}

var QualifierFlag = cli.StringFlag{
	Name:  "qualifier",
	Usage: "[optional] an additional custom version `QUALIFIER`: <version>-QUALIFIER-SNAPSHOT",
}

func snapshotRun(ctx *cli.Context) error {
	version := ctx.Args().Get(0)
	qualifier := ctx.String(QualifierFlag.Name)
	if qualifier != "" {
		qualifier = "-" + qualifier
	}
	snapshotVersion := version + qualifier + "-SNAPSHOT"
	fmt.Printf("Snapshot version: %s", snapshotVersion)
	return nil
}