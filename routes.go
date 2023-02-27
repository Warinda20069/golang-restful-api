package main

import (
	"wunkyyy-pos/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	// http://example.com/hello GET -> /Status Code
	// {"key": value} => {"greeting": "Hello Server"}
	// Routing

	// /products GET
	// /products/1 GET
	// /products/1 PATCH
	// /products/1 DELETE
	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("/:id", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)
}
