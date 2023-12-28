package middleware

import (
	"log"
	"net/http"
	"time"

	"math/rand"
)

func RandomSleep(rate float64, timeout time.Duration) func(next http.Handler) http.Handler {
	log.Printf("RandomSleep: rate=%f, timeout=%d\n", rate, timeout)
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if rand.Float64() < rate {
				log.Println("Random sleep")
				timer := time.NewTimer(timeout)
				select {
				case <-r.Context().Done():
					return
				case <-timer.C:
					timer.Stop()
					next.ServeHTTP(w, r)
				}
			} else {
				next.ServeHTTP(w, r)
			}

		}
		return http.HandlerFunc(fn)
	}
}
