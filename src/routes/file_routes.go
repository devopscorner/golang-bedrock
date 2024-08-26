// routes/file_routes.go
package routes

import (
	"github.com/devopscorner/golang-bedrock/src/controller"
	"github.com/devopscorner/golang-bedrock/src/middleware"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("file-routes")

func SetupFileRoutes(rg *gin.RouterGroup, fileController *controller.FileController) {
	api := rg.Group("/v1/files", middleware.AuthMiddleware())
	{
		api.GET("", wrapHandler("FindAll", fileController.FindAll))
		api.GET("/:id", wrapHandler("FindByID", fileController.FindByID))
		api.POST("", wrapHandler("CreateFile", fileController.CreateFile))
		api.PUT("/:id", wrapHandler("UpdateFile", fileController.UpdateFile))
		api.DELETE("/:id", wrapHandler("DeleteFile", fileController.DeleteFile))
	}
}

func wrapHandler(opName string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), opName)
		defer span.End()

		span.SetAttributes(attribute.String("http.method", c.Request.Method))
		span.SetAttributes(attribute.String("http.url", c.Request.URL.String()))

		c.Request = c.Request.WithContext(ctx)
		handler(c)

		span.SetAttributes(attribute.Int("http.status_code", c.Writer.Status()))
	}
}
