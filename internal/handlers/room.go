package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	guuid "github.com/google/uuid"
)


func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if uuid == "" {
		c.Status(400)
		return nil
	}
}


func RoomWebsocket(c *websocket.Conn) error {
	uuid := c.Params("uuid")

	if uuid == "" {
		return nil
	}

	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}