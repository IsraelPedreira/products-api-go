package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required,min=2,max=40"`
	Price       float64 `json:"price" binding:"required,number"`
	Description string  `json:"description"`
}
