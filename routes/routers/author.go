package routers


import (
    // "fmt"
	"encoding/json"

    "gopkg.in/macaron.v1"

    "github.com/mhmdryhn/gocrud/selector"
)


func AuthorHandler(ctx *macaron.Context)  {
    requestBody, err := ctx.Req.Body().String()
    // fmt.Println("Request body:\n", requestBody)

    var incomingJSON map[string]interface{}
    var response map[string]interface{}

    if err != nil {
        // fmt.Println("Error:", err)
        ctx.JSON(200, map[string]interface{} {
            "verdict": "Invalid JSON",
        })
    }

    json.Unmarshal([]byte(requestBody), &incomingJSON)

    response = selector.Execute(incomingJSON)
    // fmt.Println("Response:", response)

    ctx.JSON(200, response)
}
