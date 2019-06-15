package dbops

import (
    "fmt"
    "errors"

    "github.com/mhmdryhn/gocrud/models"
)

const (
    AUTHOR_UPDATE_KEY string = "email"
)

func UpdateAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    resp := make(map[string]interface{})
    resp, err := schemas.ValidateAuthorUpdateData(data)

    if err != nil {
        return resp, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        resp["verdict"] = "Database connection error"
        return resp, errors.New("connection_error")
    }

    if dbError := db.Where(&models.Author{Email: data[AUTHOR_UPDATE_KEY].(string)}).First(&author); dbError.Error != nil {
        fmt.Println("error:", dbError.Error, author)
        return map[string]interface{} {
            "verdict": "email address not found",
        }, nil
    }
    return data, nil
}
