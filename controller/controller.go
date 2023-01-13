package controller

import (
	"net/http"

	res "github.com/KunalAnkur/todo-app/helper"
	mongoModel "github.com/KunalAnkur/todo-app/model/mongo_model"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todoDoc := &mongoModel.Todo{}
	todos, err := todoDoc.GetAllTodo()
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseToRead(w, http.StatusOK, "Things are left to do", todos)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	createTodo := &mongoModel.Todo{}
	todo, err := createTodo.CreateTodo(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseToCreate(w, http.StatusAccepted, "Your todo's are successfully created in the database", *todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	updatedTodo := &mongoModel.Todo{}
	err := updatedTodo.UpdateTodoById(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}
	res.ResponseMessage(w, http.StatusAccepted, "Your todo has been updated successfully")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	deleteTodo := &mongoModel.Todo{}
	err := deleteTodo.DeleteTodoById(r)
	if err != nil {
		res.ResponseError(w, http.StatusInternalServerError, "Something went wrong")
	}

	res.ResponseMessage(w, http.StatusAccepted, "Your todo has been deleted successfully")
}
