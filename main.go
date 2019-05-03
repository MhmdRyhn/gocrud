package main


import (
    "os"

    "github.com/urfave/cli"

    "github.com/mhmdryhn/gocrud/routers/routes"
)


const APP_VERSION = "0.0.1"


func main ()  {
    app := cli.NewApp()
    app.Name = "Go "
    app.Usage = "Learning CRUD operation using go-macaron"
    app.Version = APP_VERSION

    app.Commands = []cli.Command {
        routes.Check,
    }

    app.Run(os.Args)
}
