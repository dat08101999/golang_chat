package main

import (
	"golang_chat/controllers"
	"golang_chat/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Post("/sendMessage/:id", controllers.UserSendMessage)
	app.Get("/user/getActive", controllers.GetListUserActive)
	app.Get("/user/getNotActive", controllers.GetListUserNotActive)

	app.Get("/user/getAllMessageOfRoom/:id/:page/:perpage", controllers.GetAllMessageOfRoom)
	app.Get("/user/getAllChatRoom/:id", controllers.GetALLRoomChat)
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		errConnect := controllers.UserConnectSocket(c.Params("id"), true)
		if errConnect == nil {
			services.SendToAll(services.MessageModel{
				From:    c.Params("id"),
				Message: "connect",
			})
		}
		/// add to list socket
		services.AddToListUserSocket(services.SocketUserModel{Socket: c, UserId: c.Params("id")})

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println(c.Cookies("Sec-WebSocket-Key"))
				/// do remove socket here
				errConnect := controllers.UserConnectSocket(c.Params("id"), false)
				if errConnect == nil {
					services.SendToAll(services.MessageModel{
						From:    c.Params("id"),
						Message: "disconnect",
					})
				}
				services.RemoveSocket(c)
				log.Println("read:", err)
				break
			}

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))
	// log.Fatal(app.Listen(":" + "3000"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
