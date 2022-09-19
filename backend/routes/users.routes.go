package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()
	var users []models.Users
	db.DB.AutoMigrate(users)
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Usuario"))
}
