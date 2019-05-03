package routers


import (
    "fmt"
	"encoding/json"

    "gopkg.in/macaron.v1"
)


func Authors(ctx *macaron.Context)  {
    fmt.Println("name: ", ctx.Params(":name"))
    requestBody, err := ctx.Req.Body().String()
    fmt.Println("Req body:", requestBody)
    var jsonresp map[string]interface{}
    json.Unmarshal([]byte(requestBody), &jsonresp)
    fmt.Println(jsonresp)
    addr := jsonresp["Addr"]
    fmt.Println(addr)
    fmt.Println("JSON to map", jsonresp["Name"], jsonresp["Email"],
        jsonresp["Addr"])
    fmt.Println(err)
    // p := ContactForm{"Rayhan", 5}
    // ctx.HTML(200, "hello")
    ctx.JSON(200, jsonresp)
}
