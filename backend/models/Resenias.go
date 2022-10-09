package models

import "gorm.io/gorm"

type Resenias struct {
	gorm.Model

	Id      int    `gorm:"not null","autoIncrement","unique_index"` // ID is the primary key
	Email   string `gorm:"not null"`                                // Title is the item title
	Name    string `gorm:"not null"`                                // Description is the item description
	Comment string `gorm:"not null"`                                // Price is the item price
}
