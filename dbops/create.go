package dbops


import (
    "github.com/mhmdryhn/gocrud/models"
)


func CreateNewAuthor(data map[string]interface{}) map[string]interface{} {
    var resp map[string]interface{}
    resp = make(map[string]interface{})
    author := models.Author{}

    if value, ok := data["name"]; ok {
        author.Name = value.(string)
    } else {
        resp["status"] = "name key is missing"
    }
    if value, ok := data["email"]; ok {
        author.Name = value.(string)
    } else {
        resp["status"] = "email key is missing"
    }
    if value, ok := data["phone"]; ok {
        author.Name = value.(string)
    }
    if value, ok := data["age"]; ok {
        author.Name = value.(uint)
    }
    if value, ok := data["address"]; ok {
        author.Name = value.(string)
    }

    db.Create(&user)

    return resp
}
