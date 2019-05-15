package main


import (
    "os"
    // "fmt"
    // "encoding/json"

    "github.com/urfave/cli"

    "github.com/mhmdryhn/gocrud/cmd"
    // "github.com/mhmdryhn/gocrud/models"
)


const APP_VERSION = "0.0.1"


func main ()  {
    // auth := &models.Author {
    //     Name: "Joe Navarro",
    //     Email: "joe@mail.com",
    //     Phone: "+102456489",
    //     Age: 52,
    //     Address: "USA",
    // }
    // e, err := json.Marshal(auth)
    // if err == nil {
    //     fmt.Println(string(e))
    // }


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
