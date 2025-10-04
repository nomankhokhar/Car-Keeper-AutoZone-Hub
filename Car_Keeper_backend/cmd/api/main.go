package main

import (
	"Car_Keeper/internal/config"
	"Car_Keeper/internal/database"
	"Car_Keeper/internal/handler"
	"Car_Keeper/internal/middleware"
	"Car_Keeper/internal/repository"
	"Car_Keeper/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database (with auto-migration and optional schema)
	db, err := database.InitDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	log.Println("Database initialized successfully")
	// Initialize repositories
	carRepo := repository.NewCarRepository(db)
	engineRepo := repository.NewEngineRepository(db)

	// Initialize services
	carService := service.NewCarService(carRepo)
	engineService := service.NewEngineService(engineRepo)

	// Initialize handlers
	carHandler := handler.NewCarHandler(carService)
	engineHandler := handler.NewEngineHandler(engineService)

	// Setup Gin router
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Car routes
		cars := v1.Group("/cars")
		{
			cars.GET("/:carid", carHandler.GetCarByID)
			cars.GET("/", carHandler.GetCarByBrand)
			cars.POST("/", carHandler.CreateCar)
			cars.PUT("/:carid", carHandler.UpdateCar)
			cars.DELETE("/:carid", carHandler.DeleteCar)
		}
		engine := v1.Group("/engines")
		{
			engine.GET("/:engineid", engineHandler.GetEngineByID)
			engine.POST("/", engineHandler.CreateEngine)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
