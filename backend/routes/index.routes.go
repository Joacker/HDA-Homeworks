package routes

import (
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()

	var users []models.Users
	db.DB.AutoMigrate(users)
	var equipos []models.Equipos
	db.DB.AutoMigrate(equipos)
	var champions []models.Champions
	db.DB.AutoMigrate(champions)

	w.Write([]byte("Hello World3"))
}
