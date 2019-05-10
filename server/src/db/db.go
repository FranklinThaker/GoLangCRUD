package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func DbConn() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Error connecting to database")
		os.Exit(3)
	} else {
		log.Println("Connected to database!")
		return db
	}
	// defer db.Close()
	return nil
}
