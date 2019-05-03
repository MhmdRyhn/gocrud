package routers


import (
    "gopkg.in/macaron.v1"
)


func Authors(ctx *macaron.Context)  {
    fmt.Println("name: ", ctx.Params(":name"))
    req, err := ctx.Req.Body().String()
    fmt.Println("Req body:", req)
    var jr map[string]interface{}
    json.Unmarshal([]byte(req), &jr)
    fmt.Println(jr)
    addr := jr["Addr"]
    fmt.Println(addr)
    fmt.Println("JSON to map", jr["Name"], jr["Email"], jr["Addr"])
    fmt.Println(err)
    // p := ContactForm{"Rayhan", 5}
    // ctx.HTML(200, "hello")
    ctx.JSON(200, jr)
}
