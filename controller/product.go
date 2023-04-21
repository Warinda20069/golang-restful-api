package controller

import (
	"net/http"
	"wunkyyy-pos/db"
	"wunkyyy-pos/dto"
	"wunkyyy-pos/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {
	// /products?search=tu
	search := ctx.Query("search")
	categoryID := ctx.Query("categoryId")
	ctx.JSON(http.StatusOK, gin.H{
		"FindAll":    "OK",
		"Search":     search,
		"CategoryID": categoryID,
	})

	var products []model.Product
	db.Conn.Find(&products)
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

	product := model.Product{
		SKU:        form.SKU,
		Name:       form.Name,
		Desc:       form.Desc,
		Price:      form.Price,
		Status:     form.Status,
		CategoryID: form.CategoryID,
		Image:      imagePath,
	}

	if err := db.Conn.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return data JSON
	ctx.JSON(http.StatusCreated, dto.ProductResponse{
		ID:         product.ID,
		SKU:        product.SKU,
		Name:       product.Name,
		Desc:       product.Desc,
		Price:      product.Price,
		Status:     product.Status,
		CategoryID: product.CategoryID,
		Image:      product.Image,
	})
}

func (p Product) Update(ctx *gin.Context) {

}

func (p Product) Delete(ctx *gin.Context) {

}
