package models

import "gorm.io/gorm"

type Champions struct {
	gorm.Model

	id_partido     int    `gorm:"not null"` // ID is the primary key
	nombre_equipo1 string `gorm:"not null"` // Title is the item title
	nombre_equipo2 string `gorm:"not null"` // Description is the item description
	categoria      string `gorm:"not null"` // Price is the item price
	resultado      int    `gorm:"not null"` // Price is the item price
}
