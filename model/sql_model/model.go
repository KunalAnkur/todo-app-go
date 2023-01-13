package sqlmodel

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KunalAnkur/todo-app/config"
	"github.com/gorilla/mux"
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

func (b *Todo) CreateTodo(r *http.Request) (*Todo, error) {
	_ = json.NewDecoder(r.Body).Decode(b)

	db.NewRecord(b)
	db.Create(&b)
	return b, nil
}

func (book Todo) GetAllTodo() ([]Todo, error) {
	var Books []Todo
	db.Find(&Books)
	return Books, nil
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db := db.Where("ID=?", Id).Find(&getTodo)
	return &getTodo, db
}

func (todo Todo) DeleteTodoById(r *http.Request) error {
	params := mux.Vars(r)
	id := params["bookId"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return err
	}
	db.Where("ID=?", ID).Delete(todo)
	return nil
}

func (todo Todo) UpdateTodoById(r *http.Request) error {
	var updateBook = &Todo{}
	_ = json.NewDecoder(r.Body).Decode(updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		return err
	}
	todoDetails, db := GetTodoById(ID)
	if updateBook.Message != "" {
		todoDetails.Message = updateBook.Message
	}
	db.Save(&todoDetails)
	return nil
}
