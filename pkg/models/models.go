package models

type Post struct {
	ID      int    `json:"id" gorm:"column:id;not null;primary key;unique"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

func (Post) TableName() string {
	return "posts"
}
