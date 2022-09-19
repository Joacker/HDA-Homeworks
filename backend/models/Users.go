package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	id_user  int    `gorm:"not null"` // ID is the primary key
	nombre   string `gorm:"not null"` // Title is the item title
	apellido string `gorm:"not null"` // Description is the item description
	email    string `gorm:"not null"` // Price is the item price
	password string `gorm:"not null"` // Price is the item price
}
