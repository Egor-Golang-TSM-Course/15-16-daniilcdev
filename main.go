package main

import (
	"context"
	"server-context/notifications"
)

func main() {
	// urls := []string{
	// 	"https://google.com",
	// 	"https://myip.com",
	// 	"https://github.com",
	// 	"https://yandex.ru",
	// 	"https://facebook.com",
	// 	"https://linkedin.com",
	// }

	// // task 1
	// cancelwebrequest.RequestUrlsWithTimeout(
	// 	context.Background(),
	// 	urls,
	// 	400*time.Millisecond)

	// // task 2
	// cancelcallstack.CallAndCancel(context.Background(), 100*time.Millisecond)

	// task 3
	notifications.SendNotifications(context.Background())
}
