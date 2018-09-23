package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DB_TYPE = "sqlite3"
	DB_PATH = "foo.db"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(DB_TYPE, DB_PATH)
	if err != nil {
		panic(err)
	}
}
