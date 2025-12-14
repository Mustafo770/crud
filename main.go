package main

// Импортируем свои папки и Gin
import (
	"github.com/mustafchon/blog-api/database" // подключение к базе (твой путь может быть другим!)
	"github.com/mustafchon/blog-api/models"   // модели данных
	"github.com/mustafchon/blog-api/routes"   // маршруты (создадим в следующем шаге)



	_ "gorm.io/driver/sqlite" // подчёркивание значит "импортируй, но не используй напрямую" — нужно для драйвера
)

// func main — это главная функция, которая запускается первой при запуске программы
func main() {
	// 1. Подключаемся к базе данных
	database.Connect()
	
	// 2. Автоматически создаём таблицы в базе по нашим моделям
	// Если таблиц нет — создаст. Если есть — проверит структуру
	database.DB.AutoMigrate(&models.Article{}, &models.Comment{}, &models.Like{})

	// 3. Настраиваем все маршруты (эндпоинты API)
	router := routes.SetupRouter()

	// 4. Запускаем сервер на порту 8080
	// Теперь по адресу http://localhost:8080 будет работать наш API
	router.Run(":8080")
}
