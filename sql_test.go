package learn_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES('4', 'mulyono')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}
	fmt.Println("Success query data")
}
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}
}
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "janu"
	password := "password1"
	// sebelum : script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	//fmt.Println("Script:", script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login with username:", username)
	} else {
		fmt.Println("Failed login")
	}
}
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	id := "4"
	username := "akbar"
	password := "password1"
	ctx := context.Background()
	script := "INSERT INTO user(id,username, password) VALUES(?,?,?)"
	_, err := db.ExecContext(ctx, script, id, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data")
}
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	email := "teko91@gmail.com"
	content := "halah opo se"
	script := "INSERT INTO comment(email, content ) VALUES(?,?)"
	result, err := db.ExecContext(ctx, script, email, content)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id:", lastInsertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO comment(email, content ) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()
	for i := 0; i < 10; i++ {
		email := "janu" + strconv.Itoa(i) + "@gmail.com"
		content := "Komen ke" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, content)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id:", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	script := "INSERT INTO comment(email, content ) VALUES(?,?)"
	// do transaction
	for i := 0; i < 10; i++ {
		email := "janu" + strconv.Itoa(i) + "@gmail.com"
		content := "Komen ke" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, content)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id:", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
