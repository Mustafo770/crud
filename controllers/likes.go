package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Mustafo770/blog-api/database"
    "github.com/Mustafo770/blog-api/models"
)

// ToggleLike godoc
// @Summary      Поставить/снять лайк
// @Description  Если лайк есть — снимает, если нет — ставит (по user_id и article_id)
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        like body models.Like true "user_id и article_id"
// @Success      200 {object} map[string]string
// @Router       /likes [post]
func ToggleLike(c *gin.Context) {
    var input models.Like
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }

    var existing models.Like
    // Ищем существующий лайк
    result := database.DB.Where("user_id = ? AND article_id = ?", input.UserID, input.ArticleID).First(&existing)

    if result.Error == nil {
        // Лайк есть — удаляем
        database.DB.Delete(&existing)
        c.JSON(http.StatusOK, gin.H{"message": "Лайк снят"})
    } else {
        // Лайка нет — создаём
        database.DB.Create(&input)
        c.JSON(http.StatusOK, gin.H{"message": "Лайк поставлен"})
    }
}