package sqlmodel

import (
	"strconv"

	"github.com/KunalAnkur/todo-app/config"
	model "github.com/KunalAnkur/todo-app/model/combine_model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	gorm.Model
	Message string `json:"message"`
}

var (
	db *gorm.DB
)

func MySqlConnect() {
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (b Todo) CreateTodo(body model.Todo) (model.Todo, error) {
	b.Message = body.Message

	db.NewRecord(b)
	db.Create(&b)
	body.ID = strconv.FormatUint(uint64(b.ID), 2)

	return body, nil
}

func (todo Todo) GetAllTodo() ([]model.Todo, error) {
	var todos []Todo
	db.Find(&todos)
	var combine_todo []model.Todo
	for i := 0; i < len(todos); i++ {
		c_todo := model.Todo{
			ID:      strconv.FormatUint(uint64(todos[i].ID), 10),
			Message: todos[i].Message,
		}
		combine_todo = append(combine_todo, c_todo)
	}
	return combine_todo, nil
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db := db.Where("ID=?", Id).Find(&getTodo)
	return &getTodo, db
}

func (todo Todo) DeleteTodoById(todoId string) error {

	id := todoId
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return err
	}
	db.Where("ID=?", ID).Delete(todo)
	return nil
}

func (todo Todo) UpdateTodoById(todoID string, combine_model model.Todo) error {

	ID, err := strconv.ParseInt(todoID, 0, 0)
	if err != nil {
		return err
	}
	todoDetails, db := GetTodoById(ID)
	if combine_model.Message != "" {
		todoDetails.Message = combine_model.Message
	}
	db.Save(&todoDetails)
	return nil
}
