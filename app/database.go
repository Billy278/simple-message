package app

import (
	"database/sql"
	"os"
	"simple-message/helper"
	"time"
)

func NewDB() *sql.DB {
	conn := os.Getenv("MYSQL_CONNECTION")
	DB, err := sql.Open("mysql", conn)
	helper.PanicIfError(err)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)

	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
