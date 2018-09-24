/*
** Initialize and connect our Database.
** GORM is our ORM of choice here.
 */

package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "os"
)

const (
	// DB_TYPE = os.Getenv["db_type"] || "sqlite3"
	// DB_PATH = os.Getenv["db_path"] || "foo.db"
	DB_TYPE = "sqlite3"
	DB_PATH = "./db/shops.db"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(DB_TYPE, DB_PATH)
	if err != nil {
		panic(err)
	}
}
