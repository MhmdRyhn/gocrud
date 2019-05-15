package dbops


import (
    "fmt"
    "errors"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/validators/schemas"
)


func GetAllAuthor() (map[string]interface{}, error) {
    db, _ := GetConnection()

    var authors []models.Author
    e := db.Find(&authors)
    fmt.Println("e:", e)
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
    resp := make(map[string]interface{})
    resp, err := schemas.ValidateAuthorFilterCondition(data)

    if err != nil {
        return resp, nil
    }

    var author models.Author
    db, err := GetConnection()

    if err != nil {
        resp["verdict"] = "Database connection error"
        return resp, errors.New("connection_error")
    }

    if dbError := db.Where(&models.Author{Email: data["email"].(string)}).First(&author); dbError.Error != nil {
        fmt.Println("error:", dbError.Error, author)
        return map[string]interface{} {
            "verdict": "record not found",
        }, nil
    }
    
    return author.ToMap(), nil

}
