package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/middleware"
	"rest-project/internal/models"
)

func SetupRoutes(r *gin.Engine) {
	// Получаем подключение к базе данных
	database := db.GetDB()

	// Роуты
	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me) // Защищённый эндпоинт
		// Пример роутов для работы с постами
		protected.GET("/posts", func(c *gin.Context) {
			// Пример получения всех постов
			var posts []models.Post
			database.Find(&posts)
			c.JSON(200, posts)
		})

		protected.POST("/posts", func(c *gin.Context) {
			// Пример создания поста
			var post models.Post
			if err := c.ShouldBindJSON(&post); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			database.Create(&post)
			c.JSON(201, post)
		})
	}
}
