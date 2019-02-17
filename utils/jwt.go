package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vpaliy/telex/model"
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

func GetUserId(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTClaims)
	return claims.ID
}

func CreateJWT(u *model.User) string {
	claims := &JWTClaims{
		u.ID,
		u.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
