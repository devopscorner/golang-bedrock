// routes/main_routes.go
package routes

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/devopscorner/golang-bedrock/src/controller"
	"github.com/devopscorner/golang-bedrock/src/middleware"
	"github.com/devopscorner/golang-bedrock/src/repository"
	"github.com/devopscorner/golang-bedrock/src/utility"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
)

var mainTracer = otel.Tracer("main-routes")

func SetupRoutes(router *gin.Engine, cfg *config.Config, s3Client *s3.Client, db *gorm.DB) {
	// Middleware
	if cfg.OtelTraceEnable == "true" {
		router.Use(otelgin.Middleware(cfg.OtelServiceName))
	}

	// Initialize Repository
	fileRepo := repository.NewFileRepository(db)

	// Initialize Controller
	fileController := controller.NewFileController(fileRepo, s3Client)

	// Routes Metrics Endpoint
	router.GET("/metrics", wrapMainHandler("PrometheusMetrics", utility.PrometheusHandler()))

	// Routes Healthcheck
	router.GET("/health", wrapMainHandler("Healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}))

	// Routes Welcome
	router.GET("/", wrapMainHandler("Welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Golang RESTful API with Amazon Bedrock!")
	}))

	// Routes Login
	router.POST("/login", wrapMainHandler("Login", controller.LoginUser))

	// Protected routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// File Routes
		SetupFileRoutes(authorized, fileController)
	}
}

func wrapMainHandler(opName string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := mainTracer.Start(c.Request.Context(), opName)
		defer span.End()

		span.SetAttributes(attribute.String("http.method", c.Request.Method))
		span.SetAttributes(attribute.String("http.url", c.Request.URL.String()))

		c.Request = c.Request.WithContext(ctx)
		handler(c)

		span.SetAttributes(attribute.Int("http.status_code", c.Writer.Status()))
	}
}
