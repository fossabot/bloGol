package database

import (
	"time"

	"github.com/bloGol/bloGol/pkg/models"
)

func (db *DataBase) GetPostByID(id int) (*models.Post, error) {
	post := models.Post{ID: id}
	err := db.Take(&post).Error
	return &post, err
}

func (db *DataBase) CreatePost(post *models.Post) (*models.Post, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	err := db.Create(post).Error
	return post, err
}

func (db *DataBase) UpdatePost(post *models.Post) (*models.Post, error) {
	post.UpdatedAt = time.Now()
	err := db.Model(post).Update(post).Error
	return post, err
}

func (db *DataBase) DeletePost(id int) error {
	return db.Delete(&models.Post{ID: id}).Error
}

func (db *DataBase) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := db.Find(&posts).Error
	return posts, err
}
