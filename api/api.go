package api

import (
	"fmt"
	"net/http"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/gramework/gramework"
)

var App *gramework.App

func New() {
	App = gramework.New()

	App.GET("/", webIndex)

	App.GET("/post/:id", webPost)

	App.GET("/editor", webEditor)

	api := App.Sub("/api")

	post := api.Sub("/post")              // '/api/post'
	post.GET("s", apiPosts)               // '/api/posts'
	post.GET("/getById", apiPostGetByID)  // '/api/post/getById'
	post.POST("/create", apiPostCreate)   // '/api/post/create'
	post.DELETE("/remove", apiPostRemove) // '/api/post/remove'
}

func Run() error {
	return App.ListenAndServe(fmt.Sprintf(
		"%s:%d",
		config.Config.GetString("server.host"), config.Config.GetInt("server.port"),
	))
}

func apiPosts(ctx *gramework.Context) {
	posts, err := getPosts(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(posts)
}

func apiPostGetByID(ctx *gramework.Context) {
	post, err := postGetByID(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}

func apiPostCreate(ctx *gramework.Context) {
	post, err := postCreate(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}

func apiPostRemove(ctx *gramework.Context) {
	if err := postRemove(ctx); err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func apiPostPublish(ctx *gramework.Context) {
	post, err := postPublish(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}
