package commands

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:   "server",
		Usage:  "Start the GitHub WebHooks server",
		Action: cmdServer,
	},
}
