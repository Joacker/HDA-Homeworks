package models

import "gorm.io/gorm"

type Equipos struct {
	gorm.Model

	id_equipo int    `gorm:"not null"` // ID is the primary key
	nombre    string `gorm:"not null"` // Title is the item title
	categoria string `gorm:"not null"` // Description is the item description
	pais      string `gorm:"not null"` // Price is the item price

}
