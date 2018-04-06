package db

import (
	"time"

	"github.com/bloGol/bloGol/pkg/models"
)

func GetPostByID(id int) (*models.Post, error) {
	post := models.Post{ID: id}
	err := DB.Take(&post).Error
	return &post, err
}

func CreatePost(post *models.Post) (*models.Post, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	err := DB.Create(post).Error
	return post, err
}

func UpdatePost(post *models.Post) (*models.Post, error) {
	post.UpdatedAt = time.Now()
	err := DB.Model(post).Update(post).Error
	return post, err
}

func DeletePost(id int) error {
	return DB.Delete(&models.Post{ID: id}).Error
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post
	err := DB.Find(&posts).Error
	return posts, err
}
