package commands

import "github.com/codegangsta/cli"

var Flags = []cli.Flag{}
var Commands = []cli.Command{
	{
		Name:   "server",
		Usage:  "Start the GitHub WebHooks server",
		Action: cmdServer,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "configuration, c",
				Value:  "/etc/github-webhooks.json",
				Usage:  "Configuration file path",
				EnvVar: "GITHUB_WEBHOOKS_CONFIG_FILE",
			},
		},
	},
}
