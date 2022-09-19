package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetChampionsHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()
	var champions []models.Champions
	db.DB.AutoMigrate(champions)
	db.DB.Find(&champions)
	json.NewEncoder(w).Encode(&champions)

}

func GetChampionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Partido"))
}
