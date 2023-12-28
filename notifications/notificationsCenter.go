package notifications

import (
	"context"
	"fmt"
	"time"
)

type Sender interface {
	Name() string
	Send(ctx context.Context, msg string)
}

type notificationCenter struct {
	senders map[string]Sender
}

func newCenter() *notificationCenter {
	return &notificationCenter{senders: make(map[string]Sender)}
}

func (nc *notificationCenter) SubscribeSender(sender Sender) {
	if _, exists := nc.senders[sender.Name()]; exists {
		fmt.Printf("sender %s already subscribed\n", sender.Name())
		return
	}

	nc.senders[sender.Name()] = sender
}

func (nc *notificationCenter) BeginSend(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})
	sendCtx, cancel := context.WithCancel(ctx)

	go func() {
		defer close(done)
		defer cancel()

		for i := 1; true; i++ {
			select {
			case <-sendCtx.Done():
				return
			default:
				msg := fmt.Sprintf("send message %d", i)
				for _, v := range nc.senders {
					go v.Send(ctx, msg)
				}

				timer := time.NewTimer(3 * time.Microsecond)
				<-timer.C
				timer.Stop()
			}
		}
	}()

	return done
}
