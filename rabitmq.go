package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
)

var rabitPort = "32791"
var url = flag.String("url", "amqp://guest:guest@127.0.0.1:"+rabitPort+"/", "Amqp url for both the publishe and subscriber")

type Rabitmq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabitmq(ctx context.Context) (*Rabitmq, error) {
	conn, err := amqp.Dial(*url)
	if err != nil {
		fmt.Println(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	return &Rabitmq{
		conn:    conn,
		channel: ch,
	}, nil
}

func (q *Rabitmq) EnqueueMessage(ctx context.Context, targetId uint64, msg []byte) error {
	_, err := q.channel.QueueDeclare(fmt.Sprintf("user_%d", targetId), true, false, false, true, nil)
	if err != nil {
		return err
	}
	return q.channel.Publish("", fmt.Sprintf("user_%d", targetId), true, false, amqp.Publishing{
		Body: msg,
	})
}
