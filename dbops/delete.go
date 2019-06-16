package dbops

import (
    "fmt"
    "errors"

    "github.com/mhmdryhn/gocrud/models"
    // "github.com/mhmdryhn/gocrud/exceptions"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)

const (
    AUTHOR_DELETE_KEY string = "email"
)


func DeleteAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    expectedData, err := schemas.ValidateAuthorDeleteCondition(data)

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

    author.Email = expectedData["email"].(string)

    // Unscoped() is for hard-delete
    dbError := db.Unscoped().Where(AUTHOR_DELETE_KEY + " = ?", expectedData["email"].(string)).Delete(&models.Author{})

    if dbError.Error == nil {
        return map[string]interface{} {
            "verdict": "delete successfull",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "internal error",
    }, errors.New("db_operation_error")
}
