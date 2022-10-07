package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
	"github.com/gorilla/mux"
)

// Obtener todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

}

// Obtener un usuario
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Get Usuario"))
	var user models.Users
	params := mux.Vars(r)
	//para extraer el id
	//id := params["Id"]
	//fmt.Println(params)
	db.DB.First(&user, params["Id"])
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// Obtener un usuario
func GetUserHandler2(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("Email")
	fmt.Println(email)
	password := r.URL.Query().Get("Password")
	fmt.Println(password)
	var user models.Users
	db.DB.First(&user, "Email = ? AND Password = ?", email, password)
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// Obtener un usuario
func GetUserHandler3(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Post Usuario"))
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	password := user.Password
	db.DB.First(&user, "Email = ? AND Password = ?", email, password)
	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario no encontrado")
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// Crear un usuario
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Post Usuario"))
	var user models.Users
	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	db.DB.First(&user, "Email = ?", email)
	if user.Id == 0 {
		createUser := db.DB.Create(&user)
		err := createUser.Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(&user)
		return
	}
	json.NewEncoder(w).Encode("Usuario ya registrado")
}

// Actualizar un usuario
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	w.Write([]byte("Update Usuario"))
}

// Eliminar un usuario
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	w.Write([]byte("Delete Usuario"))
}
