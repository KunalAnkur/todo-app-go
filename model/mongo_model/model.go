package mongomodel

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/KunalAnkur/todo-app/config"
	"github.com/gorilla/mux"
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

func (todo Todo) GetAllTodo() ([]Todo, error) {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	var todos []Todo
	if err != nil {
		return todos, err
	}
	for curr.Next(context.Background()) {

		err := curr.Decode(&todo)
		if err != nil {

			return todos, err
		}

		todos = append(todos, todo)
	}

	defer curr.Close(context.Background())
	return todos, nil
}

func (todo *Todo) CreateTodo(r *http.Request) (*Todo, error) {
	_ = json.NewDecoder(r.Body).Decode(&todo)
	_, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return todo, err
	}

	return todo, err

}

func (todo Todo) UpdateTodoById(r *http.Request) error {
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&todo)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"message": todo.Message}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (todo Todo) DeleteTodoById(r *http.Request) error {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return nil

}
