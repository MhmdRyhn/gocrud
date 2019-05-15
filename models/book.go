package models


import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)


type Book struct {
    gorm.Model
    Name string `gorm:"type:varchar(128);not null"`
    ISBN string `gorm:"type:varchar(64);unique;not null"`
    Author Author `gorm:"foreignkey:Author_id"`
    Author_id uint
}
