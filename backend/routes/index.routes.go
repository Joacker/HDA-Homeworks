package routes

import (
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()

	var users []models.Users
	db.DB.AutoMigrate(users)
	var resenias []models.Resenias
	db.DB.AutoMigrate(resenias)
	var agents []models.Agents
	db.DB.AutoMigrate(agents)

	w.Write([]byte("Hello World3"))
}
