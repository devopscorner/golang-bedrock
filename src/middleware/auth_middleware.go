// middleware/auth_middleware.go
package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/devopscorner/golang-bedrock/src/config"
	"github.com/devopscorner/golang-bedrock/src/view"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// UnexpectedSigningMethodError is a custom error type for unexpected JWT signing methods
type UnexpectedSigningMethodError struct {
	ActualMethod string
}

func (e *UnexpectedSigningMethodError) Error() string {
	return fmt.Sprintf("Unexpected signing method: %v", e.ActualMethod)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			view.ErrorAuthHeader(ctx)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			view.ErrorAuthHeader(ctx)
			return
		}

		tokenString := authHeaderParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w", &UnexpectedSigningMethodError{ActualMethod: token.Header["alg"].(string)})
			}
			return []byte(config.JWTSecret()), nil
		})

		if err != nil {
			view.ErrorInvalidToken(ctx)
			return
		}

		if !token.Valid {
			view.ErrorInvalidToken(ctx)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			view.ErrorInvalidToken(ctx)
			return
		}

		ctx.Set("Issuer", token.Claims.(jwt.MapClaims)[config.JWTIssuer()])
		ctx.Next()
	}
}

func GenerateToken(secret string, issuer string) (string, error) {
	// Set the expiration time to 1 hour from now
	expirationTime := time.Now().Add(time.Hour * 1).Unix()

	// Create the JWT claims
	claims := jwt.StandardClaims{
		Issuer:    config.JWTIssuer(),
		ExpiresAt: expirationTime,
	}

	// Create the JWT token with the claims and secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func ValidateCredentials(username string, password string) bool {
	return username == viper.GetString("JWT_AUTH_USERNAME") && password == viper.GetString("JWT_AUTH_PASSWORD")
}
