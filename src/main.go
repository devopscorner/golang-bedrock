package main

import (
	"fmt"
	"log"

	"github.com/devopscorner/golang-restfulapi-bedrock/src/config"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/driver"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/routes"
	"github.com/devopscorner/golang-restfulapi-bedrock/src/utility"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
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
		log.Fatalf("Failed to initialize S3 client: %v", err)
	}

	// Initialize Bedrock client
	err = utility.InitBedrock(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Bedrock client: %v", err)
	}

	// Initialize Prometheus metrics
	utility.InitMetrics()

	// Initialize Loki logger
	err = utility.InitLokiLogger(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Loki logger: %v", err)
	}

	// Set Gin mode
	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, cfg, s3Client, driver.DB)

	// Start the server
	port := fmt.Sprintf(":%d", cfg.AppPort)
	log.Printf("Server is running on %s", port)
	router.Run(port)
}
