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

func GetReseniaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Resenia"))
}
