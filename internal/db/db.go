package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB = new(gorm.DB)

func Open(path string) error {
	var err error
	DB, err = gorm.Open("sqlite3", path)
	if err != nil {
		return err
	}

	DB.SingularTable(false)
	DB.LogMode(true)

	return DB.DB().Ping()
}
