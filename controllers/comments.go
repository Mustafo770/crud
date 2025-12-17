package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Mustafo770/blog-api/database"
    "github.com/Mustafo770/blog-api/models"
	
    _ "gorm.io/gorm/clause"
)

// CreateComment godoc
// @Summary      Добавить комментарий
// @Description  Добавляет комментарий к статье
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment body models.Comment true "Текст и ID статьи"
// @Success      200 {object} models.Comment
// @Failure      400 {object} map[string]string
// @Router       /comments [post]
func CreateComment(c *gin.Context) {
    var comment models.Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }
    database.DB.Create(&comment)
    c.JSON(http.StatusOK, comment)
}

// GetComments godoc
// @Summary      Получить комментарии статьи
// @Description  Возвращает все комментарии к статье
// @Tags         comments
// @Produce      json
// @Param        article_id path int true "ID статьи"
// @Success      200 {array} models.Comment
// @Router       /comments/{article_id} [get]
func GetComments(c *gin.Context) {
    articleID := c.Param("article_id")
    var comments []models.Comment
    database.DB.Where("article_id = ?", articleID).Find(&comments)
    c.JSON(http.StatusOK, comments)
}

// DeleteComment godoc
// @Summary      Удалить комментарий
// @Description  Удаляет комментарий по ID
// @Tags         comments
// @Param        id path int true "ID комментария"
// @Success      200 {object} map[string]string
// @Router       /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
    id := c.Param("id")
    database.DB.Delete(&models.Comment{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Комментарий удалён"})
}
