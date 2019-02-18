package main

import (
	"github.com/labstack/echo"
	_ "github.com/vpaliy/telex/channels"
	"github.com/vpaliy/telex/db"
	"github.com/vpaliy/telex/handler"
	_ "github.com/vpaliy/telex/messages"
	"github.com/vpaliy/telex/router"
	"github.com/vpaliy/telex/rtm"
	_ "github.com/vpaliy/telex/users"
)

var (
	manager = rtm.NewChannelManager(&rtm.TestRepository{})
)

func registerHTTPHandlers(g *echo.Group, hs ...handler.Handler) {
	for _, handler := range hs {
		handler.Register(g)
	}
}

func registerRTM(e *echo.Echo) {
	e.GET("/ws", func(c echo.Context) error {
		ws := rtm.NewWebSocket(rtm.DefaultWebSocketConfig)
		client := rtm.NewClient(ws, manager)
		return client.ServeHTTP(c.Response(), c.Request())
	})
}

func main() {
	e := router.New()
	//api := e.Group("/api")

	database, err := db.New(db.CreateTestConfig())
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer database.Close()
	db.AutoMigrate(database)

	/*	registerHTTPHandlers(api,
		users.NewHandler(database),
		channels.NewHandler(database),
		messages.NewHandler(database),
	)  */
	registerRTM(e)
	go manager.Run()
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
