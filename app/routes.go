package app

import (
	"github.com/gorilla/mux"
	"github.com/ssr0016/gotoko/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
