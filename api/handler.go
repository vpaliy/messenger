package api

import (
	"github.com/labstack/echo"
)

type Handler interface {
	Register(group *echo.Group)
}

type Binder struct{}

func (r *Binder) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
