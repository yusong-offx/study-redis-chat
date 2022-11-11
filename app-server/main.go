package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/websocket/v2"
	"github.com/yusong-offx/study-redis-chat/components"
)

const timeoutLimit = time.Second * 10

func main() {
	// init
	components.RedisConnect()
	defer components.RedisClient.Close()

	// fiber start
	app := fiber.New()
	app.Use("/ws", timeout.New(components.WebsocketConnectFunc, timeoutLimit))
	app.Get("/ws/:chatroom/:id", timeout.New(websocket.New(components.ChatFunc), timeoutLimit))

	log.Fatal(app.Listen(":3000"))
}
