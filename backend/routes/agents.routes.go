package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func GetAgentsHandler(w http.ResponseWriter, r *http.Request) {
	//db.DBConnection()
	var agents []models.Agents
	//db.DB.AutoMigrate(champions)
	db.DB.Find(&agents)
	json.NewEncoder(w).Encode(&agents)

}

func GetAgentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Agents"))
}
