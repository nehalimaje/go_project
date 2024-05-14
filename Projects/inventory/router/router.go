package router

import (
	"inventory/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/inventoryapi/employee", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/inventoryapi/employee", controller.GetEmployees).Methods("GET")
	router.HandleFunc("/inventoryapi/employee/{id}", controller.GetEmployee).Methods("GET")
	router.HandleFunc("/inventoryapi/employee/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/inventoryapi/employee/{id}", controller.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/inventoryapi/employee", controller.DeleteAllEmployee).Methods("DELETE")
	return router
}
