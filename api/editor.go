package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gramework/gramework"
)

const editorForm = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Make new post</title>
  </head>
  <body>
    <form method="post" action="/publish">
      <p><b>Title:</b> <input name="title"></p>
      <p><b>Content:</b> <textarea name="content" cols="45" rows="20"></textarea></p>
      <p><button type="submit">Publish!</button></p>
    </form>
  </body>
</html>`

func handlerEditor(ctx *gramework.Context) {

	var err error
	tpl := template.New("editor")
	tpl, err = tpl.Parse(editorForm)
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
		ctx.SuccessString(
			"text/html",
			fmt.Sprintln("<b>title:</b>", title, "<br>", "<b>content:</b>", content),
		)
	default:
		ctx.Error("please, fill and title, and content fields", http.StatusBadRequest)
	}
}
