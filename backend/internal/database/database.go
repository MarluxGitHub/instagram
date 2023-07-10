package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB = nil

func Connect() *gorm.DB {

	if dbInstance != nil {
		return dbInstance
	}

	db, err := gorm.Open(mysql.Open("user:user@/mydb"), &gorm.Config{})

	if err != nil {
		panic ("failed to connect to Database")
	}

	defer func() {
		dbToClose, err := db.DB()
		if err != nil {
			panic("failed to close database connection")
		}
		dbToClose.Close()
	}()

	dbInstance = db

	return db
}