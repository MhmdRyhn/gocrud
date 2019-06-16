package cmd


import (
    "github.com/urfave/cli"

    "github.com/mhmdryhn/gocrud/routes"
    "github.com/mhmdryhn/gocrud/dbops"
)

var StartServer = cli.Command {
    Name: "serve",
    Usage: "Start web server",
    Description: "Run web server after setting up database",
    Action: startServer,
}


var SetupDB = cli.Command {
    Name: "setupdb",
    Usage: "Generate the database tables",
    Description: "Run this command to setup database",
    Action: generateTables,
}


func startServer(c *cli.Context) error {
    routes.Routes()  // inside `routes.go`
    return nil
}


func generateTables(c *cli.Context) error {
    dbops.CreateTables()  // inside `create_table.go`
    return nil
}
