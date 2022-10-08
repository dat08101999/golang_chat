package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomChat struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	List_user []string           `bson:"list_user"`
}

type ChatMessage struct {
	Id             primitive.ObjectID `bson:"_id"`
	ConversationId primitive.ObjectID `bson:"conversationId"`
	CreateAt       primitive.DateTime `bson:"createAt"`
	From           string             `bson:"from"`
	Content        string             `bson:"content"`
}
