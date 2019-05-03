package cmd


import (
    "github.com/urfave/cli"

    "github.com/mhmdryhn/gocrud/routers/routes"
)

var StartServer = cli.Command {
    Name: "startserver",
    Usage: "Start web server",
    Description: "Web server is the only thing you need to run",
    Action: startServer,
}


func startServer(c *cli.Context) error {
    routes.Routes()  // inside `routes.go`
    return nil
}
