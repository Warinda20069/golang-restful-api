package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {
	// /products?search=tu
	search := ctx.Query("search")
	categoryId := ctx.Query("categoryId")
	ctx.JSON(http.StatusOK, gin.H{
		"FindAll":    "OK",
		"Search":     search,
		"CategoryID": categoryId,
	})
}

func (p Product) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"ID": id})
}

func (p Product) Create(ctx *gin.Context) {

}

func (p Product) Update(ctx *gin.Context) {

}

func (p Product) Delete(ctx *gin.Context) {

}
