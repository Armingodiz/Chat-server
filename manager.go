package main

import (
	"context"
	"fmt"
)

type Manager struct {
	clients map[uint64]*Client

	queue *Rabitmq
}

func (m *Manager) AddClient(ctx context.Context, client *Client) error {
	m.clients[client.Id] = client
	fmt.Printf("client added successfully !", client.Id)
	return nil
}

func (m *Manager) SendMessage(ctx context.Context, targetId uint64, msg []byte) error {
	target, ok := m.clients[targetId]
	if !ok {
		err := m.queue.EnqueueMessage(ctx, targetId, msg)
		if err != nil {
			return err
		}
		fmt.Println("target not found ! message enqueued")
		return nil
	}
	return target.WriteMessage(ctx, msg)
}

func (m *Manager) DeleteClient(ctx context.Context, client *Client) error {
	delete(m.clients, client.Id)
	return nil
}
