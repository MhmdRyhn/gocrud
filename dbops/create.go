package dbops


import (
    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func CreateNewAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    author := models.Author{}
    expectedData, err := schemas.ValidateNewAuthor(data)

    if err != nil {
        return expectedData, err
    }

    db, err := GetConnection()
    defer db.Close()

    if err != nil {
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

    if dberror := db.Create(&author); dberror.Error != nil {
        return map[string]interface{} {
            "verdict": "An author with email " + author.Email + " already exists",
        }, nil
    }

    return map[string]interface{} {
        "verdict": "Author created successfully",
    }, nil
}
