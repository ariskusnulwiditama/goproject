package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupDatabase() *gorm.DB {
    db, err := gorm.Open("mysql", "root@/goproject?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }
    return db
}