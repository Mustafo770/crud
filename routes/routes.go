package routes

import (
	"github.com/Mustafo770/blog-api/controllers"
	_ "github.com/Mustafo770/blog-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет от блог API! Сервер работает 🚀",
			"status":  "ok",
			"docs":    "Swagger документация здесь: /swagger/index.html",
		})
	})


articles := r.Group("/articles")
{
    articles.POST("/", controllers.CreateArticle)
    articles.GET("/", controllers.GetArticles)
    articles.GET("/:id", controllers.GetArticle)
    articles.PUT("/:id", controllers.UpdateArticle)
    articles.DELETE("/:id", controllers.DeleteArticle)

    articles.GET("/:id/comments", controllers.GetComments)       
    articles.POST("/:id/comments", controllers.CreateComment)    
}

comments := r.Group("/comments")
{
    comments.POST("/", controllers.CreateComment)
    comments.GET("/", controllers.GetComments) 
    comments.DELETE("/:id", controllers.DeleteComment)
}

	likes := r.Group("/likes")
	{
		likes.POST("/", controllers.ToggleLike)
	}
	return r
}
