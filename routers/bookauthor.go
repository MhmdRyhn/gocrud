package routers


import (
    "fmt"
	"encoding/json"

    "gopkg.in/macaron.v1"
)


func Authors(ctx *macaron.Context)  {
    requestBody, err := ctx.Req.Body().String()

    fmt.Println("Request body:\n", requestBody)
    fmt.Println("Error:", err)

    var jsonresp map[string]interface{}
    json.Unmarshal([]byte(requestBody), &jsonresp)
    
    fmt.Println("JSON to map:\n", jsonresp["Name"], jsonresp["Email"],
        jsonresp["Addr"])

    ctx.JSON(200, jsonresp)
}
