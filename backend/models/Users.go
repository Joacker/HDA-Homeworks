package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Id_user  int    //`gorm:"not null"` // ID is the primary key
	Nombre   string //`gorm:"not null"` // Title is the item title
	Apellido string //`gorm:"not null"` // Description is the item description
	Email    string //`gorm:"not null"` // Price is the item price
	Password string //`gorm:"not null"` // Price is the item price
}
