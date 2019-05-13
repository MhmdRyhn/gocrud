package main


import (
    "os"
    // "fmt"

    "github.com/urfave/cli"

    "github.com/mhmdryhn/gocrud/cmd"
)


const APP_VERSION = "0.0.1"


func main ()  {
    app := cli.NewApp()
    app.Name = "Go CRUD"
    app.Usage = "Learning CRUD operation using go-macaron and gorm"
    app.Version = APP_VERSION

    app.Commands = []cli.Command {
        // location: web.go
        // Generate all tables <command name: setupdb>
        cmd.SetupDB,
        // location: web.go
        //  Start the web server <command name: startserver>
        cmd.StartServer,
    }

    app.Run(os.Args)
}
