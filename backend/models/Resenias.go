package models

import "gorm.io/gorm"

type Resenias struct {
	gorm.Model

	Id        int    //`gorm:"not null"` // ID is the primary key
	Id_agente int    //`gorm:"not null"` // ID is the primary key
	Resenia   string //`gorm:"not null"` // Title is the item title
	Id_user   int    //`gorm:"not null"` // Description is the item descriptions
}
