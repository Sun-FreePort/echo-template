package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Params struct {
	Username  string
	Password  string
	Database  string
	ParseTime string
}

func New(params Params) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@/%s?parseTime=%s",
		params.Username,
		params.Password,
		params.Database,
		params.ParseTime)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
