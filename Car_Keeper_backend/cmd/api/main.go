package main

import (
	"Car_Keeper/internal/config"
	"Car_Keeper/internal/database"
	"Car_Keeper/internal/handler"
	"Car_Keeper/internal/middleware"
	"Car_Keeper/internal/repository"
	"Car_Keeper/internal/service"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Start tracing provider
	tracerProvider, err := startTracing()
	if err != nil {
		log.Fatalf("Failed to start tracing: %v", err)
	}

	// Shutdown tracing provider
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown tracer provider: %v", err)
		}
	}()

	// Set global tracer provider
	otel.SetTracerProvider(tracerProvider)

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

	// Use OpenTelemetry middleware for Gin
	router.Use(otelgin.Middleware("Car-Keeper"))
	router.Use(middleware.MetricsMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
			engine.PUT("/:engineid", engineHandler.UpdateEngine)
			engine.DELETE("/:engineid", engineHandler.DeleteEngine)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server is Started on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func startTracing() (*trace.TracerProvider, error) {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("jaeger:4318"),
			otlptracehttp.WithHeaders(header),
			otlptracehttp.WithInsecure(),
		))

	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("Car-Keeper"),
			),
		),
	)

	return tracerProvider, nil
}
