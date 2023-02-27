package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct{}

func (p Product) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"FindAll": "OK"})
}

func (p Product) FindOne(ctx *gin.Context) {

}

func (p Product) Create(ctx *gin.Context) {

}

func (p Product) Update(ctx *gin.Context) {

}

func (p Product) Delete(ctx *gin.Context) {

}
