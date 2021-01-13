package main

type Manager struct {
	clients map[uint64]*Client
}

func (m *Manager) AddClient(ctx context.Context, client *Client) error {
	m.clients[client.ID] = client
	return nil
}
