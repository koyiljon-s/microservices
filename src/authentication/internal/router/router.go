package router 

import (
	"zippilot/authentication/internal/handler"
	"zippilot/authentication/internal/middleware"
	"zippilot/authentication/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize handlers
	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)
	oauthHandler := handler.NewOAuthHandler(userRepo)

	// Setup routes
	setupPublicRoutes(r, userHandler, oauthHandler)
	setupProtectedRoutes(r, userHandler)

	return r
}

func setupPublicRoutes(r *gin.Engine, userHandler *handler.UserHandler, oauthHandler *handler.OAuthHandler) {
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
}

func setupProtectedRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	api := r.Group("/api")
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		
		protected.GET("/me", userHandler.GetMe)
		protected.PUT("/me", userHandler.UpdateMe)
		protected.DELETE("/me", userHandler.DeleteMe)

		protected.GET("/users/:id", userHandler.GetUserByID)
	}
}