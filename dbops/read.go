package dbops


import (
    "fmt"
    // "errors"

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
    // totalResp := make(map[string]interface{})

    for i := 0; i < querysetLength; i++ {
        resp[i] = map[string]interface{}{
            "ID": authors[i].ID,
            "Name": authors[i].Name,
            "Email": authors[i].Email,
            "Phone": authors[i].Phone,
            "Age": authors[i].Age,
            "Address": authors[i].Address,
        }
    }

    // fmt.Println("Authors\n=========\n", authors)
    // fmt.Println(authors[0].Email, authors[0].ID, authors[1].Email, querysetLength)
    // resp["verdict"] = "Getting"
    // totalResp["QuerySet"] = resp
    // return totalResp, nil
    
    return map[string]interface{}{
        "QuerySet": resp,
        }, nil
}


func GetAuthor(data map[string]interface{}) (map[string]interface{}, error) {
    resp := make(map[string]interface{})
    resp, err := schemas.ValidateAuthorFilterCondition(data)
}
