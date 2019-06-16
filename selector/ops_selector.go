package selector


import (
    "github.com/mhmdryhn/gocrud/dbops"
)


func Execute(data map[string]interface{}) map[string]interface{} {
    resp := make(map[string]interface{})

    value1, ok1 := data["on"]
    value2, ok2 := data["apply"]

    if !ok2 {
        resp["verdict"] = "apply key is missing"
    } else if ok1 {
        if value1 == "author" {
            if value2 == "create" {
                resp, _ = dbops.CreateNewAuthor(data)
            } else if value2 == "fetch_all" {
                resp, _ = dbops.GetAllAuthor()
            } else if value2 == "fetch_one" {
                resp, _ = dbops.GetAuthor(data)
            } else if value2 == "update" {
                resp, _ = dbops.UpdateAuthor(data)
            } else if value2 == "delete" {
                resp, _ = dbops.DeleteAuthor(data)
            } else {
                resp["verdict"] = "operation not defined (apply key's value is unknown)"
            }
        } else {
            resp["verdict"] = "Invalid value for on key"
        }
    } else {
        resp["verdict"] = "Invalid request (on key is missing)"
    }

    return resp
}
