package models

import "time"

type (
	Author struct {
		ID         int    `json:"id" gorm:"column:id;type:integer;primary_key"`
		Bio        string `json:"bio" gorm:"column:bio"`
		Location   string `json:"location" gorm:"column:location"`
		Name       string `json:"name" gorm:"column:name"`
		Slug       string `json:"slug" gorm:"column:slug"`
		Visibility string `json:"visibility" gorm:"column:visibility"`
		Website    string `json:"website" gorm:"column:website"`
	}

	Post struct {
		ID          int       `json:"id" gorm:"column:id;type:integer;primary_key"`
		AuthorID    int       `json:"author_id" gorm:"column:author_id;type:integer"`
		CreatedBy   int       `json:"created_by" gorm:"column:created_by;type:integer"`
		PublishedBy int       `json:"published_by" gorm:"column:published_by;type:integer"`
		UpdatedBy   int       `json:"updated_by" gorm:"column:updated_by;type:integer"`
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
		PublishedAt time.Time `json:"published_at" gorm:"column:published_at"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
		HTML        string    `json:"html" gorm:"column:html"`
		PlainText   string    `json:"plain_text" gorm:"column:plain_text"`
		Raw         string    `json:"raw" gorm:"column:raw"`
		Slug        string    `json:"slug" gorm:"column:slug"`
		Status      string    `json:"status" gorm:"column:status"`
		Title       string    `json:"title" gorm:"column:title"`
		Visibility  string    `json:"visibility" gorm:"column:visibility"`
	}

	Tag struct {
		ID         int       `json:"id" gorm:"column:id;type:integer;primary_key"`
		CreatedBy  int       `json:"created_by" gorm:"column:created_by;type:integer"`
		UpdatedBy  int       `json:"updated_by" gorm:"column:updated_by;type:integer"`
		Hidden     bool      `json:"hidden" gorm:"column:hidden"`
		Name       string    `json:"name" gorm:"column:name"`
		Slug       string    `json:"slug" gorm:"column:slug"`
		Visibility string    `json:"visibility" gorm:"column:visibility"`
		CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	}
)

func (Post) TableName() string { return "posts" }

func (Author) TableName() string { return "authors" }

func (Tag) TableName() string { return "tags" }
