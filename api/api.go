package api

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/bloGol/bloGol/internal/db"
	"github.com/bloGol/bloGol/pkg/models"
	"github.com/gramework/gramework"
)

var App *gramework.App

func New() {
	App = gramework.New()

	App.GET("/", GetIndexPage)
	App.GET("/posts/:id", GetPostPage)

	dashboard := App.Sub("/dashboard")
	dashboard.GET("/", GetDashboardPage)
	dashboard.GET("/editor", GetDashboardEditorPage)

	api := App.Sub("/api")

	post := api.Sub("/posts")
	post.GET("/", GetPosts)
	post.POST("/", CreatePost)
	post.GET("/:id", GetPostByID)
	post.DELETE("/:id", DeletePost)
	post.PATCH("/:id", EditPost)
}

func Run() error {
	return App.ListenAndServe(fmt.Sprintf(
		"%s:%d",
		config.Config.GetString("server.host"), config.Config.GetInt("server.port"),
	))
}

func GetPosts(ctx *gramework.Context) {
	posts, err := db.GetPosts()
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(posts)
}

func CreatePost(ctx *gramework.Context) {
	post := &models.Post{
		Title:   string(ctx.FormValue("title")),
		Content: string(ctx.FormValue("content")),
	}

	var err error
	post, err = db.CreatePost(post)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}

func GetPostByID(ctx *gramework.Context) {
	routeID, err := ctx.RouteArgErr("id")
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	id, err := strconv.Atoi(routeID)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	post, err := db.GetPostByID(id)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}

func DeletePost(ctx *gramework.Context) {
	id := int(binary.BigEndian.Uint64(ctx.QueryArgs().Peek("id")))
	if err := db.DeletePost(id); err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func EditPost(ctx *gramework.Context) {

}
