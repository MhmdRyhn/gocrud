package models


import (
    "time"

    // "github.com/jinzhu/gorm"
    // _ "github.com/jinzhu/gorm/dialects/postgres"
)


type Author struct {
    Id uint64 `xorm:"pk autoincr"`
    Name string `xorm:"VARCHAR(128) NOT NULL"`
    Email string `xorm:"VARCHAR(128) NOT NULL UNIQUE"`
    Phone string `xorm:"VARCHAR(32)"`
    Age float32
    Address string `xorm:"VARCHAR(256)"`

    Created time.Time `xorm:"created"`
    Updated time.Time `xorm:"updated"`
}


type Book struct {
    Id uint64 `xorm:"pk autoincr"`
    Name string `xorm:"varchar(128) NOT NULL"`
    Isbn string `xorm:"varchar(64) NOT NULL"`
    Author *Author `xorm:"-"`

    Created time.Time `xorm:"created"`
    Updated time.Time `xorm:"updated"`
}
