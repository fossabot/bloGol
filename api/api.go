package api

import (
	"fmt"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/gramework/gramework"
)

var Router = gramework.New()

func New() {
	Router = gramework.New()

	Router.GET("/", handlerIndex)

	Router.GET("/post/:id", handlerPost)

	Router.GET("/editor", handlerEditor)

	Router.POST("/publish", handlerPublish)

}

func Run() error {
	return Router.ListenAndServe(fmt.Sprintf(
		"%s:%d",
		config.Config.GetString("server.host"), config.Config.GetInt("server.port"),
	))
}
