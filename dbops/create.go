package dbops


import (
    // "errors"

    // "github.com/jinzhu/gorm"
    // _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/mhmdryhn/gocrud/models"
    // "github.com/mhmdryhn/gocrud/dbops"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func CreateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    var resp map[string]interface{}
    // var err error
    resp = make(map[string]interface{})
    author := models.Author{}
    resp, err := schemas.ValidateNewAuthor(data)

    // if value, ok := data["name"]; ok {
    //     author.Name = value.(string)
    // } else {
    //     resp["status"] = "name key is missing"
    // }
    // if value, ok := data["email"]; ok {
    //     author.Name = value.(string)
    // } else {
    //     resp["status"] = "email key is missing"
    // }
    // if value, ok := data["phone"]; ok {
    //     author.Name = value.(string)
    // }
    // if value, ok := data["age"]; ok {
    //     author.Name = value.(uint)
    // }
    // if value, ok := data["address"]; ok {
    //     author.Name = value.(string)
    // }
    //
    // if r, ok := resp["status"]; ok {
    //     return resp
    // }

    if err != nil {
        return resp, err
    }

    author.Name = data["name"].(string)
    author.Email = data["email"].(string)
    if value, ok := data["phone"]; !ok {
        author.Phone = value.(string)
    }
    if value, ok := data["age"]; !ok {
        author.Age = value.(float32)
    }
    if value, ok := data["address"]; !ok {
        author.Address = value.(string)
    }

    db, err := GetConnection()

    if db.NewRecord(author) {
        db.Create(&author)
        if !db.NewRecord(author) {
            db.Create(&author)
            resp["verdict"] = "Author created successfully"
            return resp, nil
        }
    }

    // db.Create(&user)

    return resp, nil
}
