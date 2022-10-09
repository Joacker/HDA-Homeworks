package models

import "gorm.io/gorm"

type Resenias struct {
	gorm.Model

	Id_resenia int    `gorm:"not null","autoIncrement","unique_index"` 
	Id_user    int    `gorm:"not null"`
	Id_agent   int    `gorm:"not null"`
	Resenia    string `gorm:"not null"` // Title is the item title
}
