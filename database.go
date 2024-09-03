package learn_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/learn_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	// nah ini adalah sql pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
