package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.MkdirAll("uploads/products", 0755)

	// engine สำหรับสร้างเว็บ
	r := gin.Default()
	serveRoutes(r)
	r.Run()
}
