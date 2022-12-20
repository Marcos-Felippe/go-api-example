package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Database() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/usersdb")
	if err != nil {
		panic(err)
	}

	return db, nil
}
