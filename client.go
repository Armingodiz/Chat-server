package main

import (
	"nhooyr.io/websocket"
)

type Client struct {
	conn *websocket.Conn
}


func (c *Client) WriteMessage(ctx context.Context,message []byte) error{
  return c.conn.Write(ctx,websocket.MessageText,message)
}
