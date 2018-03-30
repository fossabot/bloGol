package api

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/bloGol/bloGol/internal/db"
	"github.com/gramework/gramework"
)

const (
	indexPage = `<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>All posts</title>
  </head>
  <body>
    {{range .}}
        <h1><a href="/post/{{.ID}}">{{.Title}}</a></h1>
        <p>{{.Content}}</p>
        <hr>
    {{else}}
        No posts, sorry
    {{end}}
  </body>
</html>`

	postPage = `<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Post {{.ID}}</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    <p>{{.Content}}</p>
  </body>
</html>`
)

func handlerIndex(ctx *gramework.Context) {
	var list []db.Post
	db.DB.Select("*").Table(db.Post{}.TableName()).Find(&list)

	var err error
	tpl := template.New("index")
	tpl, err = tpl.Parse(indexPage)
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

	tpl := template.New("post")
	tpl, err = tpl.Parse(postPage)
	if err != nil {
		ctx.Err500(err.Error())
		return
	}

	if err = tpl.Execute(ctx.Response.BodyWriter(), post); err != nil {
		ctx.Err500(err.Error())
	}
	ctx.HTML()
}
