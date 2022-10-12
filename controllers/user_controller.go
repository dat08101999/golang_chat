package controllers

import (
	"fmt"
	"golang_chat/dbquery"
	"golang_chat/models"
	"golang_chat/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserSendMessage(c *fiber.Ctx) error {

	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	c.Response().Header.Set("Accept", "application/json")
	c.Response().Header.Set("Access-Control-Allow-Credentials", "true")
	c.Response().Header.Set("Access-Control-Allow-Headers", "GET,POST, OPTIONS")
	var messegeModel services.MessageModel
	err := c.BodyParser(&messegeModel)

	if err != nil {
		return c.JSON(map[string]interface{}{
			"err": err.Error(),
		})
	}
	chatM := models.ChatMessage{
		Id:             primitive.NewObjectID(),
		From:           c.Params("id"),
		ConversationId: messegeModel.ConversationId,
		Content:        messegeModel.Message,
		CreateAt:       primitive.NewDateTimeFromTime(time.Now()),
	}
	fmt.Println(chatM)
	// go func() {
	dbquery.SaveMessage(chatM)
	services.SendMessge(messegeModel, chatM)
	// }()
	return c.JSON(map[string]interface{}{
		"Message": "SendSuccess",
	})
}

func GetListUserActive(c *fiber.Ctx) error {
	list, err := dbquery.GetListUserActive()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"Message": "Faild",
		})
	}
	return c.JSON(map[string]interface{}{
		"list": list,
	})
}
func GetListUserNotActive(c *fiber.Ctx) error {
	list, err := dbquery.GetListUserNotActive()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"Message": "Faild",
		})
	}
	return c.JSON(map[string]interface{}{
		"list": list,
	})
}

func UserConnectSocket(name string, active bool) error {
	_, err := dbquery.FindUser(name)
	if err != nil {
		errIn := dbquery.InsertUser(name, active)
		if errIn != nil {
			return errIn
		}
	} else {
		errUp := dbquery.UpdateUser(name, active)
		if errUp != nil {
			return errUp
		}
		return nil
	}
	return nil
}

func GetALLRoomChat(c *fiber.Ctx) error {
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")

	c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

	c.Response().Header.Set("Access-Control-Allow-Headers", "GET,POST, OPTIONS")
	list, err := dbquery.GetAllRoomChat(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"Message": "error",
		})
	}
	return c.JSON(map[string]interface{}{
		"chat_room": list,
	})
}

func GetAllMessageOfRoom(c *fiber.Ctx) error {
	hexId, err := primitive.ObjectIDFromHex(c.Params("id"))
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")

	c.Response().Header.Set("Access-Control-Allow-Credentials", "true")

	c.Response().Header.Set("Access-Control-Allow-Headers", "GET,POST, OPTIONS")
	page, _ := strconv.Atoi(c.Params("page"))
	perpage, _ := strconv.Atoi(c.Params("perpage"))
	list, err := dbquery.GetMessageOfRoom(hexId, page, perpage)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"Message": "error",
		})
	}
	return c.JSON(map[string]interface{}{
		"chat_room": list,
	})
}
