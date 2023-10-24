package main

import (
	"github.com/Live-Quiz-Project/Backend/internal/db"
	"github.com/Live-Quiz-Project/Backend/internal/router"
)

func main() {
	router.InitRouter()
	db.InitDB()
	router.Start(":8080")
}
