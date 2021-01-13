package main

import (
	"context"
	"encoding/json"
	"fmt"
	"nhooyr.io/websocket"
	"time"
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
        ctx = c.conn.CloseRead(context.Background())
				c.ReadMessage(ctx, manager)
				c.ConsumeMessagesFromeQueue(ctx, manager)
				return
			}
		}
	}()
	return nil
}

func (c *Client) ReadMessage(ctx context.Context, manager *Manager) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("client " + string(c.Id) + " Done listening for messages !")
				return
			default:
				_, body, err := c.conn.Read(ctx)
				if err != nil {
					fmt.Println(err)
					continue
				}
				var message map[string]interface{}
				err = json.Unmarshal(body, &message)
				targetId := uint64(message["target"].(float64))
				if err != nil {
					fmt.Println(err)
					continue
				}
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				manager.SendMessage(ctx, targetId, body)
			}
		}
	}()
	return nil
}

func (c *Client) ConsumeMessagesFromeQueue(ctx context.Context, manager *Manager) error {
	messagesChan, err := manager.queue.ConsumeMessage(ctx, c.Id)
	if err != nil {
		return err
	}
	go func() {
		select {
		case <-ctx.Done():
			return
		case msg := <-messagesChan:
			c.WriteMessage(ctx, msg)
		}
	}()
	return nil
}
