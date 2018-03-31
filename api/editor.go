package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bloGol/bloGol/internal/db"
	"github.com/gramework/gramework"
)

func handlerEditor(ctx *gramework.Context) {
	tpl, err := template.ParseFiles("web/editor.html")
	if err != nil {
		ctx.Err500(err.Error())
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), nil); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}

func handlerPublish(ctx *gramework.Context) {
	title := string(ctx.FormValue("title"))
	content := string(ctx.FormValue("content"))

	switch {
	case title != "" && content != "":
		post := db.Post{
			Title:   title,
			Content: content,
		}
		db.DB.Create(&post)

		ctx.Redirect(fmt.Sprint("/post/", post.ID), http.StatusFound)
	default:
		ctx.Error("please, fill and title, and content fields", http.StatusBadRequest)
	}
}
