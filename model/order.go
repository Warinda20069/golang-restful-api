package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name     string
	Email    string
	Tel      string
	products []OrderItem `gorm:"foreignKey:OrderID"`
}

//somchaiOrder.OrderItems =>
