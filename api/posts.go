package api

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/bloGol/bloGol/internal/db"
	"github.com/gramework/gramework"
)

func handlerIndex(ctx *gramework.Context) {
	var list []db.Post
	db.DB.Select("*").Table(db.Post{}.TableName()).Find(&list)

	tpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), list); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}

func handlerPost(ctx *gramework.Context) {
	id, err := strconv.Atoi(ctx.RouteArg("id"))
	if err != nil {
		ctx.Redirect("/", http.StatusFound)
		return
	}

	var post db.Post
	db.DB.Select("*").Where(&db.Post{ID: int64(id)}).Find(&post)

	tpl, err := template.ParseFiles("web/post.html")
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), post); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}
