package mongomodel

import (
	"context"

	"github.com/KunalAnkur/todo-app/config"
	model "github.com/KunalAnkur/todo-app/model/combine_model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Message string             `json:"message,omitempty" `
}

var collection *mongo.Collection

func MongoConnect() {
	collection = config.GetCollection()

}

func (todo Todo) GetAllTodo() ([]model.Todo, error) {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	var todos []model.Todo
	if err != nil {
		return todos, err
	}
	for curr.Next(context.Background()) {
		var combine_todo model.Todo
		err := curr.Decode(&todo)
		if err != nil {

			return todos, err
		}
		combine_todo.ID = todo.ID.String()
		combine_todo.Message = todo.Message
		todos = append(todos, combine_todo)
	}

	defer curr.Close(context.Background())
	return todos, nil
}

func (todo Todo) CreateTodo(body model.Todo) (model.Todo, error) {
	var combine_todo model.Todo
	todo.Message = body.Message
	_, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return combine_todo, err
	}
	combine_todo.ID = todo.ID.String()
	combine_todo.Message = todo.Message
	return combine_todo, err

}

func (todo Todo) UpdateTodoById(todoID string, combine_todo model.Todo) error {
	id, _ := primitive.ObjectIDFromHex(todoID)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"message": combine_todo.Message}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (todo Todo) DeleteTodoById(todoId string) error {

	id, _ := primitive.ObjectIDFromHex(todoId)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return nil

}

// type DataStore interface {
// 	GetAllTodo() ([]model.Todo, error)
// 	CreateTodo(model.Todo) (*model.Todo, error)
// 	UpdateTodoById(string, model.Todo) error
// 	DeleteTodoById(string) error
// }
