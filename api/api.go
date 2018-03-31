package api

import (
	"fmt"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/gramework/gramework"
)

var App *gramework.App

func New() {
	App = gramework.New()

	App.GET("/", handlerIndex)

	App.GET("/post/:id", handlerPost)

	App.GET("/editor", handlerEditor)

	App.POST("/publish", handlerPublish)

}

func Run() error {
	return App.ListenAndServe(fmt.Sprintf(
		"%s:%d",
		config.Config.GetString("server.host"), config.Config.GetInt("server.port"),
	))
}
