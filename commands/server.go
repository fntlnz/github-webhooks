package commands

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/fntlnz/github-webhooks/configuration"
	"github.com/fntlnz/github-webhooks/server"
)

func cmdServer(c *cli.Context) {
	config := configuration.NewConfiguration()
	err := config.ParseFile(c.String("configuration"))
	if err != nil {
		logrus.Fatalf("Invalid configuration file: %v", err)
	}
	context := &server.Context{
		config,
	}
	handler := server.LoggingMiddleware(server.NewRouter(context))
	address := config.GetAddress()
	logrus.Infof("github-webhooks listening: %s", address)
	logrus.Fatal(http.ListenAndServe(address, handler))
}
