package notifications

import (
	"context"
	"fmt"
	"time"
)

func SendNotifications(parentCtx context.Context) {
	smsSender := newSender("SMS", func(msg string) { fmt.Println(msg) })
	emailSender := newSender("E-MAIL", func(msg string) { fmt.Println(msg) })
	pushSender := newSender("PUSH", func(msg string) { fmt.Println(msg) })

	nc := newCenter()
	nc.SubscribeSender(smsSender)
	nc.SubscribeSender(pushSender)
	nc.SubscribeSender(emailSender)

	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	done := nc.BeginSend(ctx)
	<-done
	fmt.Println("notifications stopped")
}
