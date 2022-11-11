package components

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/websocket/v2"
)

var ctx = context.Background()

func rxChat(c *websocket.Conn, user_id string, chatroom string) {
	sub := RedisClient.Subscribe(ctx, chatroom)
	defer sub.Close()
	go txChat(c, sub, user_id, chatroom)
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			break
		}
		if err := c.WriteMessage(1, []byte(msg.Payload)); err != nil {
			break
		}
	}
}

func txChat(c *websocket.Conn, sub *redis.PubSub, user_id string, chatroom string) {
	defer sub.Close()
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			break
		}
		if err := RedisClient.Publish(ctx, chatroom, user_id+" : "+string(msg)).Err(); err != nil {
			break
		}
	}
}

func ChatFunc(c *websocket.Conn) {
	user_id := c.Params("id")
	chatroom := c.Params("chatroom")
	if user_id == "" || chatroom == "" {
		return
	}
	rxChat(c, user_id, chatroom)
}
