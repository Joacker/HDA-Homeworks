package main

import (
	"fmt"
	"net/http"

	"github.com/Joacker/Ayu1/backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/champions", routes.GetChampionsHandler).Methods("GET")
	r.HandleFunc("/champions/{Id}", routes.GetChampionHandler).Methods("GET")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{Id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/equipos", routes.GetEquiposHandler).Methods("GET")
	r.HandleFunc("/equpos/{Id}", routes.GetEquipoHandler).Methods("GET")

	fmt.Println("Hello World 3")
	http.ListenAndServe(":3000", r)

}
