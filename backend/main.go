package main

import (
	"log"
	"net/http"
	"os"

	"me-too-backend/database"
	"me-too-backend/handlers"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize database connection and auto-migrations
	_, err := database.InitDB()
	if err != nil {
		log.Printf("Warning: Database initialization error: %v (backend will continue to serve status)", err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())

	// Healthcheck
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "me-too-backend"})
	})

	api := r.Group("/api")
	{
		// Auth
		api.POST("/auth/login", handlers.Login)
		api.POST("/auth/register", handlers.Register)
		api.GET("/auth/me", handlers.GetMe)

		// Daily Tasks
		api.GET("/daily-tasks", handlers.GetDailyTasks)
		api.POST("/daily-tasks", handlers.CreateDailyTask)
		api.PATCH("/daily-tasks/:id/toggle", handlers.ToggleDailyTask)
		api.DELETE("/daily-tasks/:id", handlers.DeleteDailyTask)

		// Posts & Social Feed
		api.GET("/posts", handlers.GetPosts)
		api.POST("/posts", handlers.CreatePost)
		api.POST("/posts/:id/join", handlers.JoinPost)
		api.GET("/posts/:id/thread", handlers.GetPostThread)
		api.POST("/posts/:id/submissions", handlers.AddSubmission)

		// My Tasks
		api.GET("/my-tasks", handlers.GetMyTasks)
		api.POST("/my-tasks", handlers.CreateMyTask)
		api.PATCH("/my-tasks/:id/complete", handlers.CompleteMyTask)
		api.PATCH("/my-tasks/:id/progress", handlers.UpdateTaskProgress)

		// Inspirations
		api.GET("/inspirations", handlers.GetInspirations)
		api.POST("/inspirations", handlers.CreateInspiration)
		api.DELETE("/inspirations/:id", handlers.DeleteInspiration)

		// Profile
		api.GET("/profile", handlers.GetProfile)
		api.PUT("/profile", handlers.UpdateProfile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
