package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vpaliy/telex/model"
	"time"
)

type JWTUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type JWTClaims struct {
	User *JWTUser `json:"user"`
	jwt.StandardClaims
}

// TODO: fix the secret key here
func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &JWTClaims{},
		SigningKey: []byte("secret"),
	}
	return middleware.JWTWithConfig(config)
}

func GetUser(c echo.Context) *JWTUser {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*JWTClaims)
	return claims.User
}

func GetToken(c echo.Context) *jwt.Token {
	return c.Get("user").(*jwt.Token)
}

func GetUserFromToken(token string) *JWTUser {
	return nil
}

func CreateJWT(u *model.User) string {
	claims := &JWTClaims{
		&JWTUser{
			ID:       u.ID,
			Username: u.Username,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))
	return t
}
