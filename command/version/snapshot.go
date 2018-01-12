package version

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
)

func snapshotCmd() cli.Command {
	return cli.Command {
		Name: "snapshot",
		Usage: "Process a snapshot version - set version, build, test, commit and push",
		ArgsUsage: "VERSION",
		Flags: snapshotFlags(),
		Action: func(ctx *cli.Context) error {
			return snapshotRun(ctx)
		},
	}
}

func snapshotFlags() []cli.Flag {
	return []cli.Flag {
		cli.StringFlag {
			Name: "qualifier",
			Usage: "[optional] an additional custom version `QUALIFIER`: <version>-QUALIFIER-SNAPSHOT",
		},
	}
}

func snapshotRun(ctx *cli.Context) error {
	version := ctx.Args().Get(0)
	qualifier := ctx.String("qualifier")
	if qualifier != "" {
		qualifier = "-" + qualifier
	}
	snapshotVersion := version + qualifier + "-SNAPSHOT"
	fmt.Printf("Release version: %s", snapshotVersion)
	return nil
}