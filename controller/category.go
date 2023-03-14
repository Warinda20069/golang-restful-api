package controller

import (
	"errors"
	"net/http"
	"time"
	"wunkyyy-pos/db"
	"wunkyyy-pos/dto"
	"wunkyyy-pos/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Category struct{}

func (c Category) FindAll(ctx *gin.Context) {
	// ดึง categories ทั้งหมดจาก model category มา
	var categories []model.Category
	db.Conn.Find(&categories)

	// วนลูปเอาเฉพาะ data ที่ต้องการ
	var result []dto.CategoryResponse
	for _, category := range categories {
		result = append(result, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func (c Category) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	var category model.Category
	// หาข้อมูลตัวเดียวที่เจอเป็นตัวแรก เช่น id 111
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

// ยิง request ไปที่ POST /categories <<- { "name": "Flower" } (JSON)
func (c Category) Create(ctx *gin.Context) {
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := model.Category{
		Name: form.Name,
	}
	if err := db.Conn.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ส่ง ID, Name
	// ctx.JSON(http.StatusCreated, gin.H{"id": category.ID, "name": category.Name})
	ctx.JSON(http.StatusCreated, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

func (c Category) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form dto.CategoryRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category model.Category
	if err := db.Conn.First(&category, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = form.Name
	db.Conn.Save(&category)
	ctx.JSON(http.StatusOK, dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	})

}

func (c Category) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	db.Conn.Delete(&model.Category{}, id)

	ctx.JSON(http.StatusOK, gin.H{"deletedAt": time.Now()})
}
