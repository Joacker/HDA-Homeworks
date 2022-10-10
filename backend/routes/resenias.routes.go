package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetReseniasHandler(w http.ResponseWriter, r *http.Request) {
	//db.DBConnection()
	type Result_resenia struct {
		AgentName   string `json:"Agent_Name"`
		UserEmail   string `json:"User_Email"`
		AgentComent string `json:"Agent_Coment"`
	}
	//var results Result_resenia
	var resenias []models.Resenias
	//variable para los agentes
	var results Result_resenia
	json.NewDecoder(r.Body).Decode(&results)
	agent := results.AgentName
	//select * from resenias, agents where resenias.name = 'Astra' AND resenias.name = agents.name;
	//db.DB.Select("Agents.Name, Resenias.Email, Resenias.Comment").Joins("inner join Agents on Resenias.Name = Agents.Name").Where("Resenias.Name = ?", agent).Table("Resenias").Scan(&results)
	db.DB.Select("Resenias.Name, Resenias.Email, Resenias.Comment").Where("Resenias.Name = ?", agent).Find(&resenias)
	//db.DB.AutoMigrate(equipos)
	//db.DB.Find(&results)
	json.NewEncoder(w).Encode(&resenias)

}

// Crear un usuario
func PostReseniaHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Post Usuario"))
	var resenia models.Resenias
	json.NewDecoder(r.Body).Decode(&resenia)
	email := resenia.Email
	name := resenia.Name
	db.DB.Where("Email = ? AND Name = ?", email, name).First(&resenia)
	if resenia.Id == 0 {
		createResenia := db.DB.Create(&resenia)
		err := createResenia.Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(&resenia)
		return
	}
	json.NewEncoder(w).Encode("Resenia ya registrada")
}

func PutReseniaHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Put Resenia"))
	//variable de body
	var resenia models.Resenias
	json.NewDecoder(r.Body).Decode(&resenia)
	email := resenia.Email
	name := resenia.Name
	comment := resenia.Comment
	db.DB.Where("Email = ? AND Name = ?", email, name).First(&resenia)
	if resenia.Id != 0 {
		resenia.Comment = comment
		db.DB.Save(&resenia)
		json.NewEncoder(w).Encode(&resenia)
		return
	}
	json.NewEncoder(w).Encode("No se encontro Resenia")
}

func DeleteReseniaHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Get Resenia"))
	var resenia models.Resenias
	json.NewDecoder(r.Body).Decode(&resenia)
	email := resenia.Email
	name := resenia.Name
	db.DB.Where("Email = ? AND Name = ?", email, name).First(&resenia)
	if resenia.Id != 0 {
		db.DB.Delete(&resenia)
		json.NewEncoder(w).Encode(&resenia)
		return
	}
	json.NewEncoder(w).Encode("No se encontro Resenia")
}

func GetReseniaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Resenia"))
}
