package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetReseniasHandler(w http.ResponseWriter, r *http.Request) {
	//db.DBConnection()
	var resenias []models.Resenias
	//db.DB.AutoMigrate(equipos)
	db.DB.Find(&resenias)
	json.NewEncoder(w).Encode(&resenias)

}

func GetReseniaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Resenia"))
}
