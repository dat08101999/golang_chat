package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id     primitive.ObjectID   `bson:"_id"`
	Name   string               `bson:"name"`
	Active bool                 `bson:"active"`
	ChatID []primitive.ObjectID `bson:"chatID"`
}
