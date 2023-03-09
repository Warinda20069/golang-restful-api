package dto

type ProductRequest struct {
	Name string `form:"name" binding:"required"`
}
