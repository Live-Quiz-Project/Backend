package main

import (
	"github.com/Live-Quiz-Project/Backend/internal/router"
	"github.com/Live-Quiz-Project/Backend/internal/db"
)

func main() {
	router.InitRouter()
	db.InitDB()
	router.Start(":8080")
}
