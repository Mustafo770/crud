package models


import "gorm.io/gorm"

type Article struct {
    gorm.Model  
    Title       string
    Content     string
    Comments    []Comment `gorm:"foreignKey:ArticleID"` 
    Likes       []Like    `gorm:"foreignKey:ArticleID"` 
}

type Comment struct {
    gorm.Model
    Text       string
    ArticleID  uint 
}

type Like struct {
    gorm.Model
    UserID     uint 
    ArticleID  uint
}