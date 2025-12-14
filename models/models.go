package models


import "gorm.io/gorm"

type Article struct {
    gorm.Model  // Добавляет ID, CreatedAt и т.д.
    Title       string
    Content     string
    Comments    []Comment `gorm:"foreignKey:ArticleID"` // Связь с комментариями
    Likes       []Like    `gorm:"foreignKey:ArticleID"` // Связь с лайками
}

type Comment struct {
    gorm.Model
    Text       string
    ArticleID  uint // Привязка к статье
}

type Like struct {
    gorm.Model
    UserID     uint // Предполагаем, что есть пользователи (для простоты — просто ID)
    ArticleID  uint
}