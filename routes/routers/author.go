package routers


import (
    "fmt"
	"encoding/json"

    "gopkg.in/macaron.v1"

    "github.com/mhmdryhn/gocrud/factory"
)


// func CreateAuthor(ctx *macaron.Context) map[string]interface{} {
//
// }


func AuthorHandler(ctx *macaron.Context)  {
    requestBody, err := ctx.Req.Body().String()

    fmt.Println("Request body:\n", requestBody)
    fmt.Println("Error:", err)

    var incomingJSON map[string]interface{}
    var response map[string]interface{}

    json.Unmarshal([]byte(requestBody), &incomingJSON)

    response = factory.Execute(incomingJSON)

    fmt.Println("Response:", response)

    // fmt.Println("JSON to map:\n", incomingJSON["Name"], incomingJSON["Email"],
    //     incomingJSON["Addr"])

    ctx.JSON(200, response)
}
