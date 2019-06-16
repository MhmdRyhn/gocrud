package dbops

import (
    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)

const (
    AUTHOR_DELETE_KEY string = "email"
)


func DeleteAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    expectedData, err := schemas.ValidateAuthorDeleteCondition(data)

    if err != nil {
        return expectedData, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    if dbError := db.Where(&models.Author{Email: expectedData["email"].(string)}).First(&author); dbError.Error != nil {
        return map[string]interface{} {
            "verdict": "record not found for deleting",
        }, nil
    }

    // Unscoped() is for hard-delete
    dbError := db.Unscoped().Where(AUTHOR_DELETE_KEY + " = ?", expectedData["email"].(string)).Delete(&models.Author{})

    if dbError.Error == nil {
        return map[string]interface{} {
            "verdict": "delete successfull",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "internal error",
    }, dbError.Error
}
