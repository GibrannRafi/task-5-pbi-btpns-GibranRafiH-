package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/rakamin_project"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Photo{})

	DB = db
}
