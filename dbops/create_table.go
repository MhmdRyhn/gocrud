package dbops


import (
    "fmt"
    "strings"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/mhmdryhn/gocrud/models"
    "github.com/mhmdryhn/gocrud/settings"
)


func CreateTables() error {
    dbType, credentials, err := generateDBCredentials()
    if err != nil {
        fmt.Println("settings.DATABASE not found")
    }

    db, err := gorm.Open(
        dbType,
        credentials,
    )
    defer db.Close()

    if err == nil {
        db.AutoMigrate(&models.Author{}, &models.Book{})

    } else {
        fmt.Println("Error Connecting to DB")
        return err
    }

    addDBConstraints(db)
    return nil
}


func generateDBCredentials() (string, string, error) {
    database := settings.DATABASE
    dbCredentials := ""

    temp := []string{}

    for key, value := range database {
        if value != "" {
            temp = append(temp, key + "=" + value)
        }
    }
    dbCredentials = strings.Join(temp, " ")

    return settings.DB_TYPE, dbCredentials, nil
}


func addDBConstraints(db *gorm.DB) {
    // Constraint for `models.Book`
    db.Model(models.Book{}).AddForeignKey("author_id", "authors(id)", "CASCADE", "CASCADE",)
}
