package dbops

import (
    "fmt"
    // "errors"

    "github.com/mhmdryhn/gocrud/models"
    // "github.com/mhmdryhn/gocrud/exceptions"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)

const (
    AUTHOR_UPDATE_KEY string = "email"
)

func UpdateAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    expectedData, err := schemas.ValidateAuthorUpdateData(data)

    fmt.Println("Schema Validation Error:", err)

    if err != nil {
        return expectedData, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        fmt.Println("Connection error:", err)
        // expectedData["verdict"] = "Database connection error"
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    if dbError := db.Where(&models.Author{Email: data[AUTHOR_UPDATE_KEY].(string)}).First(&author); dbError.Error != nil {
        fmt.Println("error:", dbError.Error, author)
        return map[string]interface{} {
            "verdict": "email not found",
        }, nil
    }

    if value, ok := expectedData["email"]; ok {
        author.Email = value.(string)
    }
    if value, ok := expectedData["name"]; ok {
        fmt.Println("name:", value)
        author.Name = value.(string)
    }
    if value, ok := expectedData["phone"]; ok {
        author.Phone = value.(string)
    }
    if value, ok := expectedData["age"]; ok {
        author.Age = value.(float64)
    }
    if value, ok := expectedData["Address"]; ok {
        author.Address = value.(string)
    }

    fmt.Println("updated author:", author)

    dbError := db.Save(&author)

    fmt.Println("updated error:", dbError.Error)

    if dbError.Error == nil {
        return map[string]interface{} {
            "verdict": "update successfull",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "internal error",
    }, dbError.Error
}
