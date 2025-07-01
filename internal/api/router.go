package api

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce/internal/api/middleware"
	"go-ecommerce/internal/config"
	"go-ecommerce/internal/di"

	// Swagger
	_ "go-ecommerce/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Connect DB, Redis & AutoMigrate
	config.ConnectDB()
	config.AutoMigrate()
	config.ConnectRedis()

	// === Wire Dependency Injection ===
	productHandler := di.InitializeProductHandler()
	authHandler := di.InitializeAuthHandler()
	orderHandler := di.InitializeOrderHandler()

	// === ROUTES ===

	// Auth routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Public product routes
	productGroup := r.Group("/products")
	{
		productGroup.GET("", productHandler.GetAll)
	}

	// Authenticated user info route
	protected := r.Group("/me")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("", func(c *gin.Context) {
			userID := c.GetUint("user_id")
			email := c.GetString("email")
			role := c.GetString("role")
			c.JSON(200, gin.H{
				"message": "Hello, authenticated user!",
				"user_id": userID,
				"email":   email,
				"role":    role,
			})
		})
	}

	// Admin-only product routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	{
		adminRoutes.GET("/products", productHandler.GetAll)
		adminRoutes.POST("/products", productHandler.Create)
		adminRoutes.PUT("/products/:id", productHandler.Update)
		adminRoutes.DELETE("/products/:id", productHandler.Delete)
	}

	// === USER ORDER ROUTES ===
	userOrders := r.Group("/orders")
	userOrders.Use(middleware.AuthMiddleware())
	{
		userOrders.POST("", orderHandler.Create)
		userOrders.GET("", orderHandler.GetUserOrders)
		userOrders.GET("/:id", orderHandler.GetByID)
	}

	// === ADMIN ORDER ROUTES ===
	adminOrders := r.Group("/admin/orders")
	adminOrders.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	{
		adminOrders.GET("", orderHandler.GetAll)
		adminOrders.PUT("/:id", orderHandler.UpdateStatus)
		adminOrders.DELETE("/:id", orderHandler.Delete)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}
