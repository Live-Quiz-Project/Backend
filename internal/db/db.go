package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // The database driver in use.
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=DBuser dbname=postgres_db password=DBpass sslmode=disable")

	if err != nil {
		panic(err)
	}

	// Check the database connection
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database")
}
