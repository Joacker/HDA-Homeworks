package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

// Obtener todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

// Obtener un usuario
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Usuario"))
}

// Crear un usuario
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Post Usuario"))
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

// Eliminar un usuario
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Usuario"))
}
