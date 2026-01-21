// cmd/api/main.go
package main

import (
	"log"
	
	"primejobs/user-service/internal/database"
	"primejobs/user-service/internal/handler"
	"primejobs/user-service/internal/repository"
    "primejobs/user-service/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)
	oauthHandler := handler.NewOAuthHandler(userRepo)

	// Create Gin router
	r := gin.Default() 

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Public routes
	api := r.Group("/api")
	{   
        api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})

		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)

		api.GET("/oauth/google", oauthHandler.GoogleLogin)
		api.GET("/oauth/google/callback", oauthHandler.GoogleCallback)
	}
   
    
	protected := api.Group("/")
    protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", userHandler.GetMe)
		protected.PUT("/me", userHandler.UpdateMe)
		protected.DELETE("/me", userHandler.DeleteMe)
        // internal usage
		protected.GET("/users/:id", userHandler.GetUserByID)
    }

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}