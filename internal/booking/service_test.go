package booking

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"

	"github.com/artyyomn/go-ticket-booking/internal/adapters/redis"
)

func TestConcurrency(t *testing.T) {
	store := NewRedisStore(redis.NewClient("localhost:6379"))
	svc := NewService(store)

	const Goroutines = 100_000

	var (
		succcess atomic.Int64
		failure  atomic.Int64
		wg       sync.WaitGroup
	)

	wg.Add(Goroutines)
	for i := range Goroutines {
		go func(userNum int) {
			defer wg.Done()

			err := svc.Book(Booking{
				MovieId: "movie-1",
				SeatId:  "A1",
				UserId:  uuid.New().String(),
			})

			if err == nil {
				succcess.Add(1)
			} else {
				failure.Add(1)
			}
		}(i)
	}
	wg.Wait()

	if got := succcess.Load(); got != 1 {
		t.Errorf("Expected exactly 1 success, got more")
	}

	if got := failure.Load(); got != int64(Goroutines-1) {
		t.Errorf("got more failures that expected")
	}
}
