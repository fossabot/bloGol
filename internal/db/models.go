package db

type Post struct {
	ID      int64  `gorm:"column:id;not null;primary key;unique"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

func (Post) TableName() string {
	return "posts"
}
