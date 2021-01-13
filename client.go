package main

import (
	"context"
	"encoding/json"
	"fmt"
	"nhooyr.io/websocket"
)

type Client struct {
	conn *websocket.Conn
	Id   uint64
}

func (c *Client) WriteMessage(ctx context.Context, message []byte) error {
	return c.conn.Write(ctx, websocket.MessageText, message)
}

func (c *Client) Register(ctx context.Context, manager *Manager) error {
  fmt.Println("waiting for id ...")
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("client Time for register Done !")
				return
			default:
				_, body, err := c.conn.Read(ctx)
				if err != nil {
          fmt.Println(err)
					continue
				}
				var message map[string]interface{}
				err = json.Unmarshal(body, &message)
				if err != nil {
          fmt.Println(err)
					continue
				}
				id := uint64(message["id"].(float64))
				c.Id = id
				manager.AddClient(ctx, c)
				c.ReadMessage(context.Background())
				return
			}
		}
	}()
	return nil
}

func (c *Client) ReadMessage(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("client " + string(c.Id) + " Done listening for messages !")
			return nil
		default:
			_, body, err := c.conn.Read(ctx)
			if err != nil {
				return err
			}
			var message map[string]interface{}
			err = json.Unmarshal(body, &message)
			if err != nil {
				return err
			}
		}
	}
}
