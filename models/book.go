package models


type Author struct {
    Name string `gorm:"not null"`
    Email string `gorm:"primary_key"`
    Phone string
    Age float32
    Address string
}


type Book struct {
    Name string `gorm:"not null"`
    ISBN string `gorm:"not null"`
}
