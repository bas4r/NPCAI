// app.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/basarrcan/NPCAI/routes"
	"github.com/basarrcan/NPCAI/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	a.DB = services.ConnectDB()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Printf("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	}).Methods("GET")
	a.Router.HandleFunc("/api/v1/user/new", routes.NewUserHandler).Methods("POST")
}
