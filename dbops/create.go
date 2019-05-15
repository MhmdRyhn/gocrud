package dbops


import (
    "fmt"
    "errors"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func CreateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    // var resp map[string]interface{}
    resp := make(map[string]interface{})
    author := models.Author{}
    resp, err := schemas.ValidateNewAuthor(data)

    if err != nil {
        return resp, err
    }

    author.Name = data["name"].(string)
    author.Email = data["email"].(string)
    if value, ok := data["phone"]; ok {
        author.Phone = value.(string)
    }
    if value, ok := data["age"]; ok {
        author.Age = value.(float64)
    }
    if value, ok := data["address"]; ok {
        author.Address = value.(string)
    }

    db, err := GetConnection()

    if err != nil {
        resp["verdict"] = "Database connection error"
        return resp, errors.New("connection_error")
    }


    if dberror := db.Create(&author); dberror.Error != nil {
        resp["verdict"] = dberror.Error
        fmt.Println("DB insert: ", dberror.Value)
        // resp["verdict"] = "Author created successfully"
        return resp, nil
    }

    resp["verdict"] = "Author created successfully"
    return resp, nil


    // if db.NewRecord(author) {
    //     db.Create(&author)
    //     if !db.NewRecord(author) {
    //         r := db.Create(&author)
    //         fmt.Println("Created:", r)
    //         resp["verdict"] = "Author created successfully"
    //         return resp, nil
    //     }
    // }


    return resp, nil
}
