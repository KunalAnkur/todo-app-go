package helper

import (
	"encoding/json"
	"net/http"

	model "github.com/KunalAnkur/todo-app/model/combine_model"
	"github.com/KunalAnkur/todo-app/response"
)

var DATABASE_INDEX = 1

// type Response[T string | model.Todo | []model.Todo] struct {
// 	Message    string `json:"message,omitempty" `
// 	Data       T      `json:"data,omitempty" `
// 	StatusCode int    `json:"status,omitempty" `
// }

// func (res *Response[T]) Send(w *http.ResponseWriter) {
// 	var wr = *w
// 	wr.WriteHeader(res.StatusCode)

// 	wr.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(wr).Encode(res)
// }

func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	err := response.ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
		Status:     response.FAIL_STATUS,
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func ResponseMessage(w http.ResponseWriter, statusCode int, message string) {
	resData := response.ResponseMessage{
		Message:    message,
		Status:     response.SUCCESS_STATUS,
		StatusCode: statusCode,
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func ResponseToCreate(w http.ResponseWriter, statusCode int, message string, data model.Todo) {
	resData := response.ResponseToCreate{
		Message:    message,
		Data:       data,
		Status:     response.SUCCESS_STATUS,
		StatusCode: statusCode,
	}
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func ResponseToRead(w http.ResponseWriter, statusCode int, message string, data []model.Todo) {
	resData := response.ResponseToRead{
		Message:    message,
		Data:       data,
		Status:     response.SUCCESS_STATUS,
		StatusCode: statusCode,
	}
	w.WriteHeader(statusCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}
