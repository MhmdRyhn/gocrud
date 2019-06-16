package dbops

import (
    "fmt"
    "errors"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)

const (
    AUTHOR_UPDATE_KEY string = "email"
)

func UpdateAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    requiredData, err := schemas.ValidateAuthorUpdateData(data)

    if err != nil {
        return requiredData, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        requiredData["verdict"] = "Database connection error"
        return requiredData, errors.New("connection_error")
    }

    if dbError := db.Where(&models.Author{Email: data[AUTHOR_UPDATE_KEY].(string)}).First(&author); dbError.Error != nil {
        fmt.Println("error:", dbError.Error, author)
        return map[string]interface{} {
            "verdict": "email address not found",
        }, nil
    }

    if value, ok := requiredData["email"]; ok {
        author.Email = value.(string)
    }
    if value, ok := requiredData["name"]; ok {
        fmt.Println("name:", value)
        author.Name = value.(string)
    }
    if value, ok := requiredData["phone"]; ok {
        author.Phone = value.(string)
    }
    if value, ok := requiredData["age"]; ok {
        author.Age = value.(float64)
    }
    if value, ok := requiredData["Address"]; ok {
        author.Address = value.(string)
    }

    fmt.Println("updated author:", author)

    dbErrorr := db.Save(&author)

    fmt.Println("updated error:", dbErrorr.Error)

    if dbErrorr.Error == nil {
        return map[string]interface{} {
            "verdict": "update successfull",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "internal error",
    }, errors.New("db_operation_error")
}
