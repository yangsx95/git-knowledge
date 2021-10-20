package middlewares

import (
	"github.com/golang-jwt/jwt"
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
		ContextKey: "_jwt",
		SuccessHandler: func(context echo.Context) {
			token := context.Get("_jwt").(*jwt.Token)
			claims := token.Claims.(*JWTClaims)
			context.Set("_userid", claims.Userid)
			// 清除_jwt
			context.Set("_jwt", nil)
		},
	}
	return middleware.JWTWithConfig(config)
}
