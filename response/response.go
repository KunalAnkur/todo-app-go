package response

import model "github.com/KunalAnkur/todo-app/model/mongo_model"

const SUCCESS_STATUS = "Success"
const FAIL_STATUS = "Failure"

type ErrorResponse struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}
type ResponseMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}
type ResponseToRead struct {
	Message    string       `json:"message"`
	Data       []model.Todo `json:"data"`
	Status     string       `json:"status"`
	StatusCode int          `json:"status_code"`
}

type ResponseToCreate struct {
	Message    string     `json:"message"`
	Data       model.Todo `json:"data"`
	Status     string     `json:"status"`
	StatusCode int        `json:"status_code"`
}
