package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Id        int    `gorm:"not null","autoIncrement","unique_index"` // ID is the primary key
	Email     string `gorm:"not null"`                                // Title is the item title
	Password  string `gorm:"not null"`                                // Price is the item price
	Logged_in int    `gorm:"not null"`                                // Statement event
}
