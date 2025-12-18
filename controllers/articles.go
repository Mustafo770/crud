package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Mustafo770/blog-api/database"
    "github.com/Mustafo770/blog-api/models"
    "gorm.io/gorm/clause"
)

// CreateArticle godoc
// @Summary      Создать новую статью
// @Description  Создаёт новую статью в блоге
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      models.Article  true  "Данные статьи"
// @Success      200      {object}  models.Article
// @Failure      400      {object}  map[string]string
// @Router       /articles [post]
func CreateArticle(c *gin.Context) {
    var article models.Article
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }
    
  
    database.DB.Create(&article)
    c.JSON(http.StatusOK, article)
}

// GetArticles godoc
// @Summary      Получить список статей
// @Description  Возвращает список статей с пагинацией и поиском
// @Tags         articles
// @Produce      json
// @Param        page    query     int     false  "Номер страницы"     default(1)
// @Param        limit   query     int     false  "Кол-во на странице" default(10)
// @Param        search  query     string  false  "Поиск по заголовку или тексту"
// @Success      200     {array}   models.Article
// @Router       /articles [get]
func GetArticles(c *gin.Context) {
    var articles []models.Article

    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
    search := c.Query("search")

    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    offset := (page - 1) * limit  
   
    query := database.DB.Offset(offset).Limit(limit)

    if search != "" {
        searchPattern := "%" + search + "%"
        query = query.Where("title LIKE ? OR content LIKE ?", searchPattern, searchPattern)
    }

    query.Preload(clause.Associations).Find(&articles) 
    c.JSON(http.StatusOK, articles)
}
// GetArticle godoc
// @Summary      Получить одну статью
// @Description  Возвращает статью по ID с комментариями и лайками
// @Tags         articles
// @Produce      json
// @Param        id path int true "ID статьи"
// @Success      200 {object} models.Article
// @Failure      404 {object} map[string]string
// @Router       /articles/{id} [get]
func GetArticle(c *gin.Context) {
    id := c.Param("id")
    var article models.Article

    if err := database.DB.Preload(clause.Associations).First(&article, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Статья не найдена"})
        return
    }

    c.JSON(http.StatusOK, article)
}

// UpdateArticle godoc
// @Summary      Обновить статью
// @Description  Обновляет заголовок и содержание статьи по ID
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        id path int true "ID статьи"
// @Param        article body models.Article true "Новые данные"
// @Success      200 {object} models.Article
// @Failure      400,404 {object} map[string]string
// @Router       /articles/{id} [put]
func UpdateArticle(c *gin.Context) {
    id := c.Param("id")
    var article models.Article

    if err := database.DB.First(&article, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Статья не найдена"})
        return
    }

    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }

    database.DB.Save(&article)

    c.JSON(http.StatusOK, article)
}

// DeleteArticle godoc
// @Summary      Удалить статью
// @Description  Удаляет статью и все её комментарии по ID
// @Tags         articles
// @Param        id path int true "ID статьи"
// @Success      200 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /articles/{id} [delete]
func DeleteArticle(c *gin.Context) {
    id := c.Param("id")
    var article models.Article

    if err := database.DB.First(&article, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Статья не найдена"})
        return
    }

   
    database.DB.Delete(&article)

    c.JSON(http.StatusOK, gin.H{"message": "Статья успешно удалена"})
}