package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// engine สำหรับสร้างเว็บ
	r := gin.Default()
	
	serveRoutes(r)

	r.Run()
}
