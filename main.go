package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/fntlnz/github-webhooks/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "github-webhooks"
	app.Usage = "Execute commands on the server as a result of a GitHub Web Hook request."
	app.Version = "1.0.0-dev"
	app.Author = "Lorenzo Fontana"
	app.Email = "fontanalorenzo@me.com"
	app.Commands = commands.Commands
	app.Run(os.Args)
}
