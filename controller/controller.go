package controller

import (
	"net/http"

	res "github.com/KunalAnkur/todo-app/helper"
	service "github.com/KunalAnkur/todo-app/service"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := service.GetAllTodoToResponse()
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseToRead(w, http.StatusOK, "Things are left to do", todos)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo, err := service.CreateTodoToResponse(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseToCreate(w, http.StatusAccepted, "Your todo's are successfully created in the database", todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	err := service.UpdateTodoByIdToResponse(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseMessage(w, http.StatusAccepted, "Your todo has been updated successfully")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	err := service.DeleteTodoByIdToResponse(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}

	res.ResponseMessage(w, http.StatusAccepted, "Your todo has been deleted successfully")
}
