package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string
	Prince float64
}
