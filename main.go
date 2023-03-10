package main

import (
	"os"
	"wunkyyy-pos/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()

	os.MkdirAll("uploads/products", 0755)

	// engine สำหรับสร้างเว็บ
	r := gin.Default()
	serveRoutes(r)
	r.Run()
}
