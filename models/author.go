package models


import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)


type Author struct {
    gorm.Model
    Name string `gorm:"type:varchar(128);not null"`
    Email string `gorm:"type:varchar(128);unique;not null"`
    Phone string `gorm:"type:varchar(32)"`
    Age float64
    Address string `gorm:"type:varchar(256)"`
}


func (author Author) ToMap() map[string]interface{} {
    return map[string]interface{} {
        "ID": author.ID,
        "Name": author.Name,
        "Email": author.Email,
        "Phone": author.Phone,
        "Age": author.Age,
        "Address": author.Address,
    }
}


func (author Author) ToSlice() []interface{} {
    dataReturn := make([]interface{}, 0)
    data := author.ToMap()
    for key, value := range data {
        dataReturn = append(dataReturn, key, value)
    }
    return dataReturn
}
