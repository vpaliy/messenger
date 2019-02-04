package handler

import (
	"github.com/labstack/echo"
)

type Handler interface {
	Register(group *echo.Group)
}
