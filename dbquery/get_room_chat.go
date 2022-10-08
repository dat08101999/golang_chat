package dbquery

import (
	"context"
	"golang_chat/models"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllRoomChat(name string) ([]models.RoomChat, error) {
	var list []models.RoomChat
	var errSS error
	query(func(client *mongo.Client, ctx context.Context) {
		collection := getColecttion(client, "conversations")
		cur, err := collection.Find(ctx, bson.M{"list_user": name})
		if err != nil {
			errSS = err
		}
		for cur.Next(ctx) {
			var t models.RoomChat
			errDecode := cur.Decode(&t)
			if errDecode != nil {
				println(errDecode.Error())
			}
			list = append(list, t)
		}
	})
	return list, errSS
}
func GetMessageOfRoom(id primitive.ObjectID, page int, perpage int) ([]models.ChatMessage, error) {
	var list []models.ChatMessage
	var errS error
	query(func(client *mongo.Client, ctx context.Context) {
		collection := getColecttion(client, "chat_list")
		// findOptions := options.Find()
		// findOptions.SetSort(bson.D{{"createAt", -1}})
		// cur, err := collection.Find(ctx, bson.M{"conversationId": id}, findOptions)
		agPage, errs := New(collection).Context(ctx).Limit(int64(perpage)).Page(int64(page)).Sort("createAt", -1).Aggregate(bson.M{
			"$match": bson.M{"conversationId": id},
		})

		if errs != nil {
			println(errs.Error())
			errS = errs
			// return list, errS
			return
		}
		for _, raw := range agPage.Data {
			var product models.ChatMessage
			if marshallErr := bson.Unmarshal(raw, &product); marshallErr == nil {
				list = append(list, product)
			}

		}
		// for cur.Next(ctx) {
		// 	var t models.ChatMessage
		// 	errDecode := cur.Decode(&t)
		// 	if errDecode != nil {
		// 		println(errDecode.Error())
		// 	}
		// 	list = append(list, t)
		// }
	})
	return list, errS
}
func SaveMessage(model models.ChatMessage) error {
	var err error
	query(func(client *mongo.Client, ctx context.Context) {
		collection := getColecttion(client, "chat_list")
		_, errs := collection.InsertOne(ctx, model)
		if errs != nil {
			err = errs
			println(errs.Error())
		}
	})
	return err
}
