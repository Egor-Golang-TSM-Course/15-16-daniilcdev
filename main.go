package main

import (
	"context"
	cancelwebrequest "server-context/cancelWebRequest"
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

	// task 1
	cancelwebrequest.RequestUrlsWithTimeout(
		context.Background(),
		urls,
		400*time.Millisecond)
}
