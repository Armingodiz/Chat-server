package main


import (
  "context"
  "fmt"
)
type Manager struct {
	clients map[uint64]*Client
}

func (m *Manager) AddClient(ctx context.Context, client *Client) error {
	m.clients[client.Id] = client
  fmt.Println("client added successfully !")
	return nil
}
