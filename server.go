package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/vpaliy/telex/db"
	"github.com/vpaliy/telex/handler"
	"github.com/vpaliy/telex/users"
)

func registerHandlers(g *echo.Group, hs ...*handler.Handler) {
	for _, handler := range hs {
		handler.Register(g)
	}
}

func createEcho() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	return e
}

func main() {
	e := createEcho()
	api := e.Group("/api")

	database, err := db.New(db.CreateTestConfig())
	if err != nil {
		e.Logger.Fatal(error)
	}
	database.AutoMigrate()

	registerHandlers(api, nil)
	e.Logger.Fatal(e.Start(":5000"))
}
