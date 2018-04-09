package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/bloGol/bloGol/internal/database"
	"github.com/bloGol/bloGol/pkg/models"
	"github.com/gramework/gramework"
	log "github.com/kirillDanshin/dlog"
)

type App struct{ *gramework.App }

var API *App

func New() *App {
	app := gramework.New()

	app.GET("/", GetIndexPage)
	app.GET("/posts/:id", GetPostPage)

	dashboard := app.Sub("/dashboard")
	dashboard.GET("/", GetDashboardPage)
	dashboard.GET("/editor", GetDashboardEditorPage)
	dashboard.GET("/editor/:id", GetDashboardEditorPage)

	api := app.Sub("/api")

	post := api.Sub("/posts")
	post.GET("/", GetPosts)
	post.POST("/", CreatePost)
	post.GET("/:id", GetPostByID)
	post.DELETE("/:id", DeletePost)
	post.PATCH("/:id", EditPost)

	return &App{app}
}

func (app *App) Run(cfg *config.Configuration) error {
	return app.ListenAndServe(fmt.Sprintf(
		"%s:%d",
		cfg.GetString("server.host"), cfg.GetInt("server.port"),
	))
}

func GetPosts(ctx *gramework.Context) {
	posts, err := database.DB.GetPosts()
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(posts)
}

func CreatePost(ctx *gramework.Context) {
	log.D(ctx.RequestCtx)
	post := new(models.Post)
	err := ctx.UnJSON(post)
	if err != nil {
		ctx.BadRequest(err)
		return
	}

	post, err = database.DB.CreatePost(post)
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

	post, err := database.DB.GetPostByID(id)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}

func DeletePost(ctx *gramework.Context) {
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

	if err := database.DB.DeletePost(id); err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
}

func EditPost(ctx *gramework.Context) {
	log.D(ctx)
	post := new(models.Post)
	err := ctx.UnJSON(post)
	if err != nil {
		ctx.BadRequest(err)
		return
	}

	routeID, err := ctx.RouteArgErr("id")
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	if _, err = strconv.Atoi(routeID); err != nil {
		ctx.Err500(err.Error())
		return
	}

	post, err = database.DB.UpdatePost(post)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	post, err = database.DB.GetPostByID(post.ID)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	ctx.JSON(post)
}
