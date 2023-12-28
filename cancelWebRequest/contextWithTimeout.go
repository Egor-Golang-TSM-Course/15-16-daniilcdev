package cancelwebrequest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func RequestUrlsWithTimeout(parentCtx context.Context, urls []string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go get(ctx, url, &wg)
	}

	wg.Wait()
}

func get(ctx context.Context, url string, wg *sync.WaitGroup) {
	receive := make(chan []byte)

	go func() {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

		if err != nil {
			fmt.Printf("%s - failed to create request, err %v\n", url, err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("%s - failed to send request, err %v\n", url, err)
			return
		}
		defer resp.Body.Close()

		p := []byte{}
		_, err = resp.Body.Read(p)
		if err != nil {
			fmt.Printf("%s - failed to read response, err %v\n", url, err)
			return
		}

		receive <- p
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Printf("%s - failed to handle request, err %v\n", url, err)
		}
	case <-receive:
		fmt.Printf("%s - response received\n", url)
	}

	wg.Done()
}
