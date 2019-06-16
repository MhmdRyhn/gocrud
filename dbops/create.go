package dbops


import (
    "fmt"
    // "errors"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func CreateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    // var resp map[string]interface{}
    // resp := make(map[string]interface{})
    author := models.Author{}
    expectedData, err := schemas.ValidateNewAuthor(data)

    if err != nil {
        return expectedData, err
    }

    db, err := GetConnection()

    if err != nil {
        fmt.Println("Connection error:", err)
        // expectedData["verdict"] = "Database connection error"
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    author.Name = expectedData["name"].(string)
    author.Email = expectedData["email"].(string)
    if value, ok := expectedData["phone"]; ok {
        author.Phone = value.(string)
    }
    if value, ok := expectedData["age"]; ok {
        author.Age = value.(float64)
    }
    if value, ok := expectedData["address"]; ok {
        author.Address = value.(string)
    }

    // db, err := GetConnection()
    //
    // if err != nil {
    //     resp["verdict"] = "Database connection error"
    //     return resp, errors.New("connection_error")
    // }


    if dberror := db.Create(&author); dberror.Error != nil {
        // resp["verdict"] = dberror.Error
        // fmt.Println("DB insert: ", dberror.Value)
        // return resp, nil
        return map[string]interface{} {
            "verdict": dberror.Error,
        }, nil
    }

    // resp["verdict"] = "Author created successfully"
    return map[string]interface{} {
        "verdict": "Author created successfully",
    }, nil


    // if db.NewRecord(author) {
    //     db.Create(&author)
    //     if !db.NewRecord(author) {
    //         r := db.Create(&author)
    //         fmt.Println("Created:", r)
    //         resp["verdict"] = "Author created successfully"
    //         return resp, nil
    //     }
    // }


    // return resp, nil
}
