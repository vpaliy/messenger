package main

import (
	"github.com/labstack/echo"
	"github.com/vpaliy/telex/api"
	"github.com/vpaliy/telex/db"
	"github.com/vpaliy/telex/di"
	"github.com/vpaliy/telex/router"
	"github.com/vpaliy/telex/rtm"
	"github.com/vpaliy/telex/utils"
)

func registerHTTPHandlers(g *echo.Group, hs ...api.Handler) {
	for _, handler := range hs {
		handler.Register(g)
	}
}

func registerRTM(e *echo.Echo, dispatcher rtm.Dispatcher) {
	e.GET("/ws", func(c echo.Context) error {
		ws := rtm.NewWebSocket(rtm.DefaultWebSocketConfig)
		client := rtm.NewClient(ws, dispatcher, utils.GetToken(c))
		return client.ServeHTTP(c.Response(), c.Request())
	})
}

func main() {
	e := router.New()
	api := e.Group("/api")
	database, err := db.New(db.CreateTestConfig())
	if err != nil {
		e.Logger.Fatal(err)
	}

	defer database.Close()
	db.AutoMigrate(database)

	registerHTTPHandlers(api,
		di.InitializeUserHandler(database),
		di.InitializeChannelHandler(database),
		di.InitializeMessageHandler(database),
	)

	dispatcher := di.InitializeDispatcher(database)
	registerRTM(e, dispatcher)

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
