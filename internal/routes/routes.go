package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {

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
	}

}
