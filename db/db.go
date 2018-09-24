/*
** Initialize and connect our Database.
** GORM is our ORM of choice here.
 */

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "os"
)

const (
	// 	DB_TYPE = os.Getenv("DB_TYPE")
	// 	DB_PATH = os.Getenv("DB_PATH")
	DB_TYPE = "sqlite3"
	DB_PATH = "./db/shops.db"
)

// var DB_TYPE string = os.Getenv("DB_TYPE")
// var DB_PATH string = os.Getenv("DB_PATH")

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(DB_TYPE, DB_PATH)
	if err != nil {
		panic(err)
	}
}
