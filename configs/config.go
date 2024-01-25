package configs

import (
	"rest-api/structs"

	"github.com/jinzhu/gorm"
)

// DBConnect create connection to database
func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Album{})
	db.AutoMigrate(structs.Person{})
	return db
}
