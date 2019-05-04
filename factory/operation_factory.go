package factory


import (
    "fmt"

    "github.com/mhmdryhn/gocrud/dbops"
)


func Execute(data map[string]interface{}) map[string]interface{} {
    value1, ok1 := data["on"]
    value2, ok2 := data["apply"]

    var resp map[string]interface{}

    if ok1 && ok2 {
        if value1 == "author" {
            if value2 == "create" {
                fmt.Println("Received the correct thing")
                resp = dbops.CreateNewAuthor(data)
            }
        }
    } else if ok1 {

    }

    return resp
}
