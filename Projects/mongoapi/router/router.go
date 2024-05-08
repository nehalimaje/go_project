package router

import (
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/getmovie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreatMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
