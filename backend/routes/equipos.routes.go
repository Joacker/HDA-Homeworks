package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetEquiposHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()
	var equipos []models.Equipos
	db.DB.AutoMigrate(equipos)
	db.DB.Find(&equipos)
	json.NewEncoder(w).Encode(&equipos)

}

func GetEquipoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Item"))
}
