package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // The database driver in use.
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "DBuser"
	password = "DBpass"
	dbname   = "lqp_db"
)

func InitDB() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlconn)

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
