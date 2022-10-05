package models

import "gorm.io/gorm"

type Resenias struct {
	gorm.Model

	Id       int `gorm:"not null","autoIncrement","unique_index"` // ID is the primary key
	Id_user  int
	Id_agent int
	Resenia  string //`gorm:"not null"` // Title is the item title
	Users    Users  `gorm:"foreignKey:Id_user"`
	Agents   Agents `gorm:"foreignKey:Id_agent"`
}
