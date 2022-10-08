package dbquery

import (
	"context"
	"errors"
	"golang_chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetListUserActive() ([]models.UserModel, error) {
	var errG error
	var userList []models.UserModel
	query(func(client *mongo.Client, ctx context.Context) {
		colection := getColecttion(client, "user")
		cur, err := colection.Find(ctx, bson.M{
			"active": true,
		})
		if err != nil {
			errG = err
		}
		for cur.Next(ctx) {
			var t models.UserModel
			err := cur.Decode(&t)
			if err != nil {
				println("error")
			}
			userList = append(userList, t)
		}
	})
	return userList, errG
}

func GetListUserNotActive() ([]models.UserModel, error) {
	var errG error
	var userList []models.UserModel
	query(func(client *mongo.Client, ctx context.Context) {
		colection := getColecttion(client, "user")
		cur, err := colection.Find(ctx, bson.M{
			"active": false,
		})
		if err != nil {
			errG = err
		}
		for cur.Next(ctx) {
			var t models.UserModel
			err := cur.Decode(&t)
			if err != nil {
				println("error")
			}
			userList = append(userList, t)
		}
	})
	return userList, errG
}

func InsertUser(name string, active bool) error {
	var errorG error
	query(func(client *mongo.Client, ctx context.Context) {
		colection := getColecttion(client, "user")
		_, err := colection.InsertOne(ctx, models.UserModel{
			Id:     primitive.NewObjectID(),
			Name:   name,
			Active: active,
		})
		if err != nil {
			errorG = err
		}
	})
	return errorG
}

func UpdateUser(name string, active bool) error {
	var errorG error
	query(func(client *mongo.Client, ctx context.Context) {
		colection := getColecttion(client, "user")
		_, err := colection.UpdateOne(ctx, bson.D{{Key: "name", Value: name}},
			bson.D{{Key: "$set", Value: bson.D{
				{"active", active},
			}}})
		if err != nil {
			errorG = err
		}
	})
	return errorG
}
func FindUser(name string) (models.UserModel, error) {
	var isFindUser error
	var userModel models.UserModel
	query(func(client *mongo.Client, ctx context.Context) {
		colection := getColecttion(client, "user")
		errFindOne := colection.FindOne(ctx, bson.M{"name": name}).Decode(&userModel)
		if errFindOne != nil {
			isFindUser = errors.New("User name does not exist")
		}
	})
	return userModel, isFindUser
}
