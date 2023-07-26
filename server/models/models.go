package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Task     string             `json:"task" bson:"task"`
	Status   bool               `json:"status" bson:"status"`



}
