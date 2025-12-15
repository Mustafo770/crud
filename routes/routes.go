package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
    _ "github.com/Mustafo770/blog-api/docs" 
    "github.com/Mustafo770/blog-api/controllers"
)

// SetupRouter — главная функция, которая настраивает все маршруты
// Возвращает готовый сервер Gin
func SetupRouter() *gin.Engine {
	// Создаём сервер с стандартными настройками (логи и т.д.)
	r := gin.Default()

	// Добавляем Swagger по адресу /swagger/index.html
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Тестовый маршрут на главную страницу (тот, что ты сейчас видишь)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет от блог API! Сервер работает 🚀",
			"status":  "ok",
			"docs":    "Swagger документация здесь: /swagger/index.html",
		})
	})

	
	// Группа маршрутов для статей
	articles := r.Group("/articles")
	{
		articles.POST("/", controllers.CreateArticle) // Создать статью
		articles.GET("/", controllers.GetArticles)    // Список статей с пагинацией и поиском
		
	}

	return r
}
