package notifications

import (
	"context"
	"fmt"
)

func newSender(name string, send func(string)) Sender {
	return &sender{name: name, send: send}
}

type sender struct {
	name string
	send func(string)
}

func (s *sender) Send(ctx context.Context, msg string) {
	select {
	case <-ctx.Done():
		fmt.Printf("canceled, not sent - [%s] %s\n", s.name, msg)
	default:
		s.send(fmt.Sprintf("[%s] %s", s.name, msg))
	}
}

func (s *sender) Name() string {
	return s.name
}
