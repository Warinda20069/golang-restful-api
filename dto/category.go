package dto

type CategoryRequest struct {
	Name string `json:"name" binding:"required"` // มาจาก name ตรงส่วน Body / binding required ถ้าไม่มี name จะเกิด error
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
