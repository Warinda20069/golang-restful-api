package controller

import (
	"net/http"
	"wunkyyy-pos/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var form dto.ProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath := "./uploads/products/" + uuid.New().String()
	ctx.SaveUploadedFile(image, imagePath)

	ctx.JSON(http.StatusOK, gin.H{"name": form.Name, "ImagePath": imagePath})
}

func (p Product) Update(ctx *gin.Context) {

}

func (p Product) Delete(ctx *gin.Context) {

}
