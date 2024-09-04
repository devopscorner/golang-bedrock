// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/devopscorner/golang-bedrock/src/driver"
	"github.com/devopscorner/golang-bedrock/src/routes"
	"github.com/devopscorner/golang-bedrock/src/utility"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("✗ Failed to load configuration: %v", err)
	}

	// Initialize logger
	config.InitLogger()

	// Connect to database
	driver.ConnectDatabase()

	// Initialize tracer
	cleanup := utility.InitTracer(cfg)
	defer cleanup()

	// Initialize S3 client
	s3Client, err := utility.InitS3Client(cfg)
	if err != nil {
		log.Fatalf("✗ Failed to initialize S3 client: %v", err)
	}

	// Initialize Bedrock client
	err = utility.InitBedrock(cfg)
	if err != nil {
		log.Fatalf("✗ Failed to initialize Bedrock client: %v", err)
	}

	// Initialize Prometheus metrics
	utility.InitMetrics()

	// Initialize Loki logger
	err = utility.InitLokiLogger(cfg)
	if err != nil {
		log.Printf("❗ Warning: Failed to initialize Loki logger: %v", err)
		log.Println("✔ Continuing without Loki logging...")
	} else {
		log.Println("✔ Loki logger initialized successfully")
	}

	// Set Gin mode
	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Setup CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	routes.SetupRoutes(router, cfg, s3Client, driver.DB)

	// Start the server
	port := fmt.Sprintf(":%d", cfg.AppPort)
	log.Printf("✔ Server is running on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("✗ Failed to start server: %v", err)
	}
}
