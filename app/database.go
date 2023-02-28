package app

import (
	"database/sql"
	"simple-message/helper"
	"time"
)

func NewDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/simple_message?parseTime=true")
	helper.PanicIfError(err)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)

	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
