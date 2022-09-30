package models

import "gorm.io/gorm"

type Resenias struct {
	gorm.Model

	Id      int    `gorm:"not null","autoIncrement"` // ID is the primary key
	Resenia string //`gorm:"not null"` // Title is the item title
	Users   Users  `gorm:"foreignKey:Id"`
	Agents  Agents `gorm:"foreignKey:Id"`
}
