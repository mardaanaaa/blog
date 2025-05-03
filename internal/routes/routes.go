package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/middleware"
	"rest-project/internal/models"
	"rest-project/internal/repository"
	"rest-project/internal/services"
	"strconv"
)

func SetupRoutes(r *gin.Engine) {
	// DB
	database := db.GetDB()
	postRepo := repository.NewPostRepository(database)
	postService := services.NewPostService(postRepo)

	// Auth routes
	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	// Protected routes
	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)

		// GET all posts
		protected.GET("/posts", func(c *gin.Context) {
			posts, err := postService.GetAllPosts()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, posts)
		})

		// GET post by ID
		protected.GET("/posts/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
				return
			}
			post, err := postService.GetPostByID(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
				return
			}
			c.JSON(http.StatusOK, post)
		})

		// CREATE new post
		protected.POST("/posts", func(c *gin.Context) {
			var post models.Post
			if err := c.ShouldBindJSON(&post); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			createdPost, err := postService.CreatePost(post)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, createdPost)
		})

		// UPDATE post
		protected.PUT("/posts/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
				return
			}
			var post models.Post
			if err := c.ShouldBindJSON(&post); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := postService.UpdatePost(id, post); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
		})

		// DELETE post
		protected.DELETE("/posts/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
				return
			}
			if err := postService.DeletePost(id); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
		})
	}
}
