package api

import (
	"html/template"
	"net/http"

	"github.com/gramework/gramework"
)

func webEditor(ctx *gramework.Context) {
	tpl, err := template.ParseFiles("web/editor.html")
	if err != nil {
		ctx.Err500(err.Error())
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), nil); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}

func webIndex(ctx *gramework.Context) {
	posts, err := getPosts(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	tpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), posts); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}

func webPost(ctx *gramework.Context) {
	postID, err := ctx.RouteArgErr("id")
	if err != nil {
		ctx.Redirect("/", http.StatusFound)
		return
	}

	ctx.Request.URI().QueryArgs().Add("id", postID)
	post, err := postGetByID(ctx)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

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
