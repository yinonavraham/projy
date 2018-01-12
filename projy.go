package main

import (
	"os"
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/yinonavraham/projy/command/version"
)

func main() {
	//TODO Check for updates
	err := createApp().Run(os.Args)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func createApp() *cli.App {
	app := cli.NewApp()
	configureBasics(app)
	configureCommands(app)
	return app
}

func configureBasics(app *cli.App) {
	app.Name = "projy"
	app.Usage = "A utility for common tasks in a development project lifecycle"
	app.Version = "0.1.0"
	app.EnableBashCompletion = true
}

func configureCommands(app *cli.App) {
	app.Commands = []cli.Command {
		version.Command(),
	}
}
