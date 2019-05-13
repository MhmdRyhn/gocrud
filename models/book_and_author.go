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


type Book struct {
    gorm.Model
    Name string `gorm:"type:varchar(128);not null"`
    ISBN string `gorm:"type:varchar(64);not null"`
    Author Author `gorm:"foreignkey:Author_id"`
    Author_id uint
}
