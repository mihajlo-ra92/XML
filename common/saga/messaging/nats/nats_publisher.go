package nats

import (
	"fmt"

	saga "github.com/mihajlo-ra92/XML/common/saga/messaging"
	"github.com/nats-io/nats.go"
)

type NewPCoublisher struct {
	Conn    *nats.Conn
	Subject string
}

type Publisher struct {
	conn    *nats.EncodedConn
	subject string
}

func NewNATSPublisher(host, port, user, password, subject string) (saga.Publisher, error) {
	conn, err := getConnection(host, port, user, password)
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Publisher{
		conn:    encConn,
		subject: subject,
	}, nil
}

func NewPCublisher(host, port, user, password, subject string) (*nats.Conn, error) {
	url := fmt.Sprintf("nats://%s:%s@%s:%s", user, password, host, port)
	connection, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (p *Publisher) Publish(message interface{}) error {
	err := p.conn.Publish(p.subject, message)
	if err != nil {
		return err
	}
	return nil
}
