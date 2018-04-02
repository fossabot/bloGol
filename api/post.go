package api

import (
	"github.com/bloGol/bloGol/internal/db"
	"github.com/bloGol/bloGol/pkg/models"
	"github.com/gramework/gramework"
	log "github.com/kirillDanshin/dlog"
)

func postGetByID(ctx *gramework.Context) (*models.Post, error) {
	log.Ln(ctx.QueryArgs().String())
	post := models.Post{
		ID: ctx.QueryArgs().GetUintOrZero("id"),
	}

	err := db.DB.Take(&post).Error
	return &post, err
}

func postCreate(ctx *gramework.Context) (*models.Post, error) {
	log.D(ctx.Request)
	post := models.Post{
		Title:   string(ctx.QueryArgs().Peek("title")),
		Content: string(ctx.QueryArgs().Peek("content")),
	}

	err := db.DB.Create(&post).Error
	return &post, err
}

func postRemove(ctx *gramework.Context) error {
	post := models.Post{
		ID: ctx.QueryArgs().GetUintOrZero("id"),
	}

	return db.DB.Delete(&post).Error
}

func getPosts(ctx *gramework.Context) ([]models.Post, error) {
	var posts []models.Post
	err := db.DB.Find(&posts).Error
	return posts, err
}
