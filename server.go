package main

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/db"
	"github.com/vpaliy/telex/handler"
	"github.com/vpaliy/telex/router"
	"github.com/vpaliy/telex/users"
)

func registerHandlers(g *echo.Group, hs ...handler.Handler) {
	for _, handler := range hs {
		handler.Register(g)
	}
}

func main() {
	e := router.New()
	api := e.Group("/api")

	database, err := db.New(db.CreateTestConfig())
	if err != nil {
		e.Logger.Fatal(err)
	}
	database.AutoMigrate()

	registerHandlers(api, users.NewHandler(database))
	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
