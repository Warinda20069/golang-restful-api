package main

import (
	"wunkyyy-pos/controller"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {
	// http://example.com/hello GET -> /Status Code
	// {"key": value} => {"greeting": "Hello Server"}
	// Routing
	
	categoryController := controller.Category{}
	categoryGroup := r.Group("/categories")
	categoryGroup.GET("", categoryController.FindAll)
	categoryGroup.GET("/:id", categoryController.FindOne)
	categoryGroup.POST("/:id", categoryController.Create)
	categoryGroup.PATCH("/:id", categoryController.Update)
	categoryGroup.DELETE("/:id", categoryController.Delete)

	productController := controller.Product{}
	productGroup := r.Group("/products")
	productGroup.GET("", productController.FindAll)
	productGroup.GET("/:id", productController.FindOne)
	productGroup.POST("/:id", productController.Create)
	productGroup.PATCH("/:id", productController.Update)
	productGroup.DELETE("/:id", productController.Delete)
}
