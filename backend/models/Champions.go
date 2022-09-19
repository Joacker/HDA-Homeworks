package models

import "gorm.io/gorm"

type Champions struct {
	gorm.Model

	Id_partido     int    `gorm:"not null"` // ID is the primary key
	Nombre_equipo1 string `gorm:"not null"` // Title is the item title
	Nombre_equipo2 string //`gorm:"not null"` // Description is the item description
	Categoria      string //`gorm:"not null"` // Price is the item price
	Resultado      int    //`gorm:"not null"` // Price is the item price
}
