// controller/login_controller.go
package controller

import (
	"github.com/devopscorner/golang-bedrock/src/middleware"
	"github.com/devopscorner/golang-bedrock/src/utility"
	"github.com/devopscorner/golang-bedrock/src/view"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

var loginTracer = otel.Tracer("login-controller")

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var loginRequest LoginRequest

func LoginUser(c *gin.Context) {
	ctxTrace, span := loginTracer.Start(c.Request.Context(), "LoginUser")
	defer span.End()

	if err := c.BindJSON(&loginRequest); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Failed to bind JSON", err)
		view.ErrorBadRequest(c, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		utility.RecordError(ctxTrace, err)
		utility.LogError(c, "Validation failed", err)
		view.ErrorBadRequest(c, err)
		return
	}

	validCred := middleware.ValidateCredentials(loginRequest.Username, loginRequest.Password)
	if !validCred {
		utility.RecordError(ctxTrace, nil)
		utility.LogError(c, "Invalid credentials", nil)
		view.ErrorInvalidCredentials(c)
		return
	}

	// token, err := middleware.GenerateToken(loginRequest.Username)
	token, err := middleware.GenerateToken(viper.GetString("JWT_SECRET"), viper.GetString("JWT_AUTH_USERNAME"))

	if err != nil {
		utility.RecordError(ctxTrace, nil)
		utility.LogError(c, "Failed to generate token", err)
		view.ErrorInternalServer(c, err)
		return
	}

	view.LoginToken(c, token)
}
