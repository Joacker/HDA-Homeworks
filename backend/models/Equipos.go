package models

import "gorm.io/gorm"

type Equipos struct {
	gorm.Model

	Id_equipo int    //`gorm:"not null"` // ID is the primary key
	Nombre    string //`gorm:"not null"` // Title is the item title
	Categoria string //`gorm:"not null"` // Description is the item description
	Pais      string //`gorm:"not null"` // Price is the item price

}
