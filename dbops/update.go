package dbops

import (

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)

const (
    AUTHOR_UPDATE_KEY string = "email"
)

func UpdateAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    expectedData, err := schemas.ValidateAuthorUpdateData(data)

    if err != nil {
        return expectedData, nil
    }

    var author models.Author
    db, err := GetConnection()
    defer db.Close()

    if err != nil {
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    if dbError := db.Where(&models.Author{Email: data[AUTHOR_UPDATE_KEY].(string)}).First(&author); dbError.Error != nil {
        return map[string]interface{} {
            "verdict": "email not found",
        }, nil
    }

    author.Email = expectedData["email"].(string)
    if value, ok := expectedData["name"]; ok {
        author.Name = value.(string)
    }
    if value, ok := expectedData["phone"]; ok {
        author.Phone = value.(string)
    }
    if value, ok := expectedData["age"]; ok {
        author.Age = value.(float64)
    }
    if value, ok := expectedData["address"]; ok {
        author.Address = value.(string)
    }

    dbError := db.Save(&author)

    if dbError.Error == nil {
        return map[string]interface{} {
            "verdict": "update successfull",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "internal operation error",
    }, dbError.Error
}
