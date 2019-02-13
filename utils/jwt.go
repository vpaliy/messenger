package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

type JWTClaims struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte("secret"),
	}
	return middleware.JWTWithConfig(config)
}

func GetJWTClaims(c echo.Context) *JWTClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*JWTClaims)
}

func CreateJWT(id uint, username string) string {
	claims := &JWTClaims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
