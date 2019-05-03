package routes


import (
    "fmt"

    "github.com/urfave/cli"
    "gopkg.in/macaron.v1"
)

var runServer = cli.Command {
    Name: "runserver",
    Usage: "Start web server",
    Description: "Web server is the only thing you need to run
        It'll take care of everything"
    Action: runServer,
}

func runServer (c *cli.Context) error {
    m := macaron.Classic()
    m.Use(macaron.Renderer())
    
    m.Post("/author/list/", Authors)
}
