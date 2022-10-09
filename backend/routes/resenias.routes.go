package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
)

func GetReseniasHandler(w http.ResponseWriter, r *http.Request) {
	//db.DBConnection()
	type Result struct {
		AgentName string `json:"Agent_Name"`
	}
	var results Result
	json.NewDecoder(r.Body).Decode(&results)
	agent := results.AgentName
	//select * from resenias, agents where resenias.name = 'Astra' AND resenias.name = agents.name;
	db.DB.Select("*").Where("Resenia.Name = ? AND Resenia.Name = Agents.Name", agent).Table("Resenia", "Agents")
	//db.DB.AutoMigrate(equipos)
	db.DB.Find(&results)
	json.NewEncoder(w).Encode(&results)

}

func GetReseniaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Resenia"))
}
