package learn_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/learn_golang_database")
	if err != nil {
		panic(err)
	}
	// use db
	defer db.Close()
}
