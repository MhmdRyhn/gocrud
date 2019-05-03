package routes


import (
    "gopkg.in/macaron.v1"

    "github.com/mhmdryhn/gocrud/routers"
)


func Routes() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    m.Post("/author/list/", routers.Authors)  // inside `bookauthor.go`

    m.Run()
}
