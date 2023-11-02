package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // The database driver in use.
)

var DB *sql.DB

var GormDB *gorm.DB

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

func InitGormDB() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	GormDB = db

	fmt.Println("Connected to the database with GORM")
}
