package main

import (
	"log"
	"os"
	"wunkyyy-pos/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDB()
	db.Migrate()

	os.MkdirAll("uploads/products", 0755)

	
	r := gin.Default()
	serveRoutes(r)
	r.Run()
}
