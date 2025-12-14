package database
import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
    if err != nil {
        panic("Не удалось подключиться к БД!")
    }
    DB = db
}