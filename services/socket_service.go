package services

import (
	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

type SocketUserModel struct {
	Socket *websocket.Conn
	UserId string
}

type MessageModel struct {
	To             string             `bson:"to"`
	From           string             `bson:"from"`
	Message        string             `bson:"message"`
	ConversationId primitive.ObjectID `bson:"conversationId"`
}

var SocketUserList []SocketUserModel

func AddToListUserSocket(socketUserModel SocketUserModel) {
	SocketUserList = append(SocketUserList, socketUserModel)
}

func RemoveSocket(sc *websocket.Conn) {
	for i := 0; i < len(SocketUserList); i++ {
		if SocketUserList[i].Socket == sc {
			SocketUserList = append(SocketUserList[:i], SocketUserList[i+1:]...)
			break
		}
	}
}
func findSocket(userID []string, callback func(sc *websocket.Conn)) {
	for _, d := range SocketUserList {
		if slices.Contains(userID, d.UserId) {
			callback(d.Socket)
		}
	}
}
func SendMessge(messegeModel MessageModel) {
	var list []string
	list = append(list, messegeModel.To, messegeModel.From)
	findSocket(list, func(sc *websocket.Conn) {
		sc.WriteJSON(messegeModel)
	})
}

func SendToAll(mess MessageModel) {
	for i := 0; i < len(SocketUserList); i++ {
		SocketUserList[i].Socket.WriteJSON(mess)
	}
}
