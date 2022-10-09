package models

import "gorm.io/gorm"

type Agents struct {
	gorm.Model

	Id          int    `gorm:"not null","autoIncrement","unique_index"` // ID is the primary key
	Name        string `gorm:"not null"`                                // Title is the item title
	Role        string `gorm:"not null"`                                // Description is the item description
	Descripcion string `gorm:"not null"`                                // Price is the item price
	Skill1      string `gorm:"not null"`                                // Price is the item price
	Skill2      string `gorm:"not null"`                                // Price is the item price
	Skill3      string `gorm:"not null"`                                // Price is the item price
	Ulti        string `gorm:"not null"`                                // Price is the item price
	Url_image   string `gorm:"not null"`                                // Price is the item price
}
