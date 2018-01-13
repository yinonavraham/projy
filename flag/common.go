package flag

import "gopkg.in/urfave/cli.v1"

var ForceFlag = cli.BoolFlag {
	Name: "force",
	Usage: "Force an operation, ignore preconditions",
}

var DryRunFlag = cli.BoolFlag {
	Name: "dry-run",
	Usage: "Perform a 'dry-run', does not execute actual actions",
}

var SkipTestsFlag = cli.BoolFlag {
	Name: "skip-tests",
	Usage: "Skip running tests",
}

var SkipPushFlag = cli.BoolFlag {
	Name: "skip-push",
	Usage: "Skip pushing changes to remote repository",
}

var SkipTagFlag = cli.BoolFlag {
	Name: "skip-tag",
	Usage: "Skip tagging the last commit",
}
