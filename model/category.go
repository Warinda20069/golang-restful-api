package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(100);not null"` // Name ห้ามซ้ำ/ชื่อเป็นอักขระไม่เกิน100/ต้องระบุ name ทุกครั้ง
}
