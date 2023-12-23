package main

import (
	"context"
	cancelcallstack "server-context/cancelCallStack"
	cancelwebrequest "server-context/cancelWebRequest"
	chiserver "server-context/chiServer"
	"server-context/notifications"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://myip.com",
		"https://github.com",
		"https://yandex.ru",
		"https://facebook.com",
		"https://linkedin.com",
	}

	ctx := context.Background()
	// task 1
	cancelwebrequest.RequestUrlsWithTimeout(
		ctx,
		urls,
		400*time.Millisecond)

	// task 2
	cancelcallstack.CallAndCancel(ctx, 100*time.Millisecond)

	// task 3
	notifications.SendNotifications(ctx)

	// task 4
	chiserver.StartServer()
}
