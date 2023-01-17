package service

import (
	"encoding/json"
	"net/http"

	"github.com/KunalAnkur/todo-app-go/helper"
	model "github.com/KunalAnkur/todo-app-go/model/combine_model"
	mongomodel "github.com/KunalAnkur/todo-app-go/model/mongo_model"
	mysqlmodel "github.com/KunalAnkur/todo-app-go/model/sql_model"
	"github.com/gorilla/mux"
)

type DataStore interface {
	GetAllTodo() ([]model.Todo, error)
	CreateTodo(model.Todo) (model.Todo, error)
	UpdateTodoById(string, model.Todo) error
	DeleteTodoById(string) error
}

// type databases interface {
// 	mongomodel.Todo | mysqlmodel.Todo
// }

func getAllTodosFromDataStore(d DataStore) ([]model.Todo, error) {
	return d.GetAllTodo()
}

func GetAllTodoToResponse() ([]model.Todo, error) {
	switch helper.DATABASE_INDEX {
	case 1:
		mongoModel := mongomodel.Todo{}
		return mongoModel.GetAllTodo()
		// return getAllTodosFromDataStore(mongoModel)
	case 2:
		mysqlModel := mysqlmodel.Todo{}
		return mysqlModel.GetAllTodo()
		// return getAllTodosFromDataStore(mysqlModel)
	}

	return nil, nil
}

func CreateTodoToResponse(r *http.Request) (model.Todo, error) {
	var combine_todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&combine_todo)
	switch helper.DATABASE_INDEX {
	case 1:

		mongoModel := mongomodel.Todo{}
		return mongoModel.CreateTodo(combine_todo)
		// return getAllTodosFromDataStore(mongoModel)
	case 2:
		mysqlModel := mysqlmodel.Todo{}
		return mysqlModel.CreateTodo(combine_todo)
		// return getAllTodosFromDataStore(mysqlModel)
	}

	return combine_todo, nil
}

func UpdateTodoByIdToResponse(r *http.Request) error {
	var combine_todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&combine_todo)
	params := mux.Vars(r)
	switch helper.DATABASE_INDEX {
	case 1:
		return mongomodel.Todo{}.UpdateTodoById(params["id"], combine_todo)
	case 2:
		return mysqlmodel.Todo{}.UpdateTodoById(params["id"], combine_todo)
	}

	return nil
}

func DeleteTodoByIdToResponse(r *http.Request) error {
	params := mux.Vars(r)
	switch helper.DATABASE_INDEX {
	case 1:
		return mongomodel.Todo{}.DeleteTodoById(params["id"])
	case 2:
		return mysqlmodel.Todo{}.DeleteTodoById(params["id"])
	}

	return nil
}

// func createTodoInDataStore(d DataStore) (*model.Todo, error) {
// 	return d.CreateTodo()
// }

// func updateTodoInDataStore(d DataStore) error {
// 	return d.UpdateTodoById()
// }

// func deleteTodoInDataStore(d DataStore) error {
// 	return d.DeleteTodoById()
// }
