package factory


import (
    "fmt"

    "github.com/mhmdryhn/gocrud/dbops"
)


func Execute(data map[string]interface{}) map[string]interface{} {
    value1, ok1 := data["on"]
    value2, ok2 := data["apply"]

    // var resp map[string]interface{}
    resp := make(map[string]interface{})

    if ok1 && ok2 {
        if value1 == "author" {
            if value2 == "create" {
                fmt.Println("Received the correct thing")
                resp, _ = dbops.CreateNewAuthor(data)
            } else if value2 == "fetch_all" {
                resp, _ = dbops.GetAllAuthor()
            } else if value2 == "fetch_one" {
                resp, _ = dbops.GetAuthor(data)
            }
        }
    } else {
        resp["verdict"] = "Invalid request"
    }

    return resp
}
