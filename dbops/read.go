package dbops


import (
    "fmt"
    // "errors"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func GetAllAuthor() (map[string]interface{}, error) {
    // db, _ := GetConnection()

    var authors []models.Author
    db, err := GetConnection()

    if err != nil {
        fmt.Println("Connection error:", err)
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    dbError := db.Find(&authors)
    fmt.Println("e:", dbError)
    if dbError.Error != nil {
        return map[string]interface{} {
            "verdict": "record not found",
        }, nil
    }
    
    querysetLength := len(authors)
    resp := make([]map[string]interface{}, querysetLength)

    for i := 0; i < querysetLength; i++ {
        resp[i] = authors[i].ToMap()
    }

    return map[string]interface{}{
        "QuerySet": resp,
        }, nil
}


func GetAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    expectedData, err := schemas.ValidateAuthorFilterCondition(data)

    fmt.Println("Schema Validation Error:", err)

    if err != nil {
        return expectedData, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        fmt.Println("Connection error:", err)
        return map[string]interface{} {
            "verdict": "Database connection error",
        }, err
    }

    if dbError := db.Where(&models.Author{Email: expectedData["email"].(string)}).First(&author); dbError.Error != nil {
        fmt.Println("error:", dbError.Error, author)
        return map[string]interface{} {
            "verdict": "record not found",
        }, nil
    }

    return author.ToMap(), nil

}
