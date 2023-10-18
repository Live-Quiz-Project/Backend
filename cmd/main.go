package main

import "github.com/Live-Quiz-Project/Backend/internal/router"

func main() {
	router.InitRouter()
	router.Start(":8080")
}
