package routes

import (
	"github.com/Mustafo770/blog-api/controllers"
	_ "github.com/Mustafo770/blog-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	// Комментарии привязаны к статье — делаем вложенный маршрут
articles := r.Group("/articles")
{
    articles.POST("/", controllers.CreateArticle)
    articles.GET("/", controllers.GetArticles)
    articles.GET("/:id", controllers.GetArticle)
    articles.PUT("/:id", controllers.UpdateArticle)
    articles.DELETE("/:id", controllers.DeleteArticle)

    // Вложенная группа: комментарии к конкретной статье
    articles.GET("/:id/comments", controllers.GetComments)       // Список комментариев к статье
    articles.POST("/:id/comments", controllers.CreateComment)    // Создать комментарий к статье
}

// Отдельный маршрут для удаления комментария по его ID
comments := r.Group("/comments")
{
    comments.POST("/", controllers.CreateComment)
    comments.GET("/", controllers.GetComments) // Теперь GET /comments?article_id=123
    comments.DELETE("/:id", controllers.DeleteComment)
}

	likes := r.Group("/likes")
	{
		likes.POST("/", controllers.ToggleLike)
	}
	return r
}
