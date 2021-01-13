package main

import (
	"context"
	"github.com/streadway/amqp"
)

var url = flag.String("url", "amqp://guest:guest@127.0.0.1:55005/", "Amqp url for both the publishe and subscriber")

type Rabitmq struct {
	*amqp.Connection
	*amqp.Channel
}

func NewRabitmq(ctx context.Context) (*rabitmq, error) {
	conn, err := amqp.Dial(*url)
	if err != nil {
		fmt.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	return &rabitmq{
		conn:    conn,
		channel: ch,
	}, nil
}

func (q *rabitmq) EnqueueMessage(ctx context.Context, targetId uint64, msg []byte) error {
	_, err := q.channel.QueueDeclare(fmt.Sprintf("user_%d", targetId), true, false, false, true, nil)
	if err != nil {
		return err
	}
	return q.channel.Publish("", fmt.Sprintf("user_%d", targetId), true, false, false, amqp.Publishing{
		Body: msg,
	})
}
