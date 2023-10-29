package main

import (
	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/router"
)

func main() {
	router.InitRouter()
	db.InitDB()
	db.InitGormDB()
	router.Start(":8080")
}
