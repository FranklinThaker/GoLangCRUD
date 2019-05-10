package migrations

import (
	"CRUD_DEMO/server/src/db"
	"CRUD_DEMO/server/src/models"
)

func Migrate() {
	db := db.DbConn()
	users := models.Student{}
	db.AutoMigrate(&users)
	defer db.Close()
}
