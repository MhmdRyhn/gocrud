package routes


import (
    "gopkg.in/macaron.v1"

    "github.com/mhmdryhn/gocrud/routes/routers"
)


func Routes() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    m.Post("/author/", routers.AuthorHandler)
    // m.Post("/book/", routers.Books)

    m.Run()
}
