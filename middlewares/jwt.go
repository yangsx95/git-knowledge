package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTClaims struct {
	Userid string `json:"userid"`
	jwt.StandardClaims
}

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte(secret),
	}
	return middleware.JWTWithConfig(config)
}
