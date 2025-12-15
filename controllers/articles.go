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
    
    // Читаем JSON из тела запроса (то, что прислал пользователь)
    if err := c.ShouldBindJSON(&article); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
        return
    }
    
  
    database.DB.Create(&article)
    
    // Возвращаем созданную статью клиенту
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

    // Читаем параметры из URL: ?page=1&limit=10&search=go
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
    search := c.Query("search")

    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    offset := (page - 1) * limit  
   
    query := database.DB.Offset(offset).Limit(limit)

    // Если есть поиск — ищем по заголовку или содержимому
    if search != "" {
        searchPattern := "%" + search + "%"
        query = query.Where("title LIKE ? OR content LIKE ?", searchPattern, searchPattern)
    }

    // Загружаем статьи + связанные комментарии и лайки
    query.Preload(clause.Associations).Find(&articles)

    // Возвращаем список
    c.JSON(http.StatusOK, articles)
}