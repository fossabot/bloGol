package db

import (
	"github.com/bloGol/bloGol/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Open(path string) error {
	var err error
	DB, err = gorm.Open("sqlite3", path)
	if err != nil {
		return err
	}

	DB.SingularTable(false)
	DB.LogMode(true)

	if err = DB.DB().Ping(); err != nil {
		return err
	}

	DB.CreateTable(&models.Post{})
	DB.CreateTable(&models.Author{})
	DB.CreateTable(&models.Tag{})

	return nil
}
