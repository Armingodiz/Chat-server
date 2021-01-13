package main

import (
	"context"
	"errors"
	"fmt"
)

type Manager struct {
	clients map[uint64]*Client
}

func (m *Manager) AddClient(ctx context.Context, client *Client) error {
	m.clients[client.Id] = client
	fmt.Printf("client added successfully !", client.Id)
	return nil
}

func (m *Manager) SendMessage(ctx context.Context, targetId uint64, msg []byte) error {
	target, ok := m.clients[targetId]
	if !ok {
		return errors.New("target not found !")
	}
	return target.WriteMessage(ctx, msg)
}
