package commands
import (
	"github.com/codegangsta/cli"
	"github.com/fntlnz/github-webhooks/server"
	"log"
	"net/http"
)

func cmdServer(c *cli.Context) {
	log.Fatal(http.ListenAndServe(":8080", server.NewRouter()))
}