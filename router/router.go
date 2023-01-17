package router

import (
	"github.com/KunalAnkur/todo-app-go/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/showtodos", controller.GetAllTodo).Methods("GET")
	router.HandleFunc("/create", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/delete/{id}", controller.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/update/{id}", controller.UpdateTodo).Methods("PUT")

	return router
}
