package database

import (
	// "fmt"
	"fmt"

	"github.com/garciasa/machinedirectory/server/storage"

	"github.com/jinzhu/gorm"
	// Driver for mysql backend
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	// Driver for sqlite backed
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Database Type for store database structure
type Database struct{ Db *gorm.DB }

// New returns a database
func New(user, password, dbname, driver string) (Database, error) {
	var d Database
	var err error
	var connect string
	if driver == "mysql" {
		connect = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
		d.Db, err = gorm.Open("mysql", connect)
	} else {
		d.Db, err = gorm.Open("sqlite3", "./db.sqlite")
	}

	if err != nil {
		return d, err
	}

	return d, nil
}

// CreateStructure create scheme in database
func (d *Database) CreateStructure() {
	d.Db.AutoMigrate(&storage.Item{})
}
