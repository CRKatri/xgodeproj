package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/CRKatri/xgodeproj/command"
)

// GlobalFlags that uses cli flags
var GlobalFlags = []cli.Flag{}

// Commands have command list
var Commands = []cli.Command{
	{
		Name:   "show",
		Usage:  "Prints section names or each section information",
		Action: command.CmdShow,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "section",
				Value: "",
				Usage: "section name for pbxproj",
			},
			cli.StringFlag{
				Name:  "project",
				Value: "",
				Usage: "xcode project name",
			},
		},
	},
}

// CommandNotFound error
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
