package version

import (
	"gopkg.in/urfave/cli.v1"
)

func Command() cli.Command {
	return cli.Command {
		Name:        "version",
		Usage:       "Project version management commands",
		Subcommands: subCommands(),
	}
}

func subCommands() []cli.Command {
	return []cli.Command {
		releaseCmd(),
		rcCmd(),
		snapshotCmd(),
		currentCmd(),
	}
}