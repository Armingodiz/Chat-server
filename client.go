package main

import (
	"context"
  //"net/http"
	"nhooyr.io/websocket"
)

type Client struct {
	conn *websocket.Conn
	Id   uint64
}

func (c *Client) WriteMessage(ctx context.Context, message []byte) error {
	return c.conn.Write(ctx, websocket.MessageText, message)
}
