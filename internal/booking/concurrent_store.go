package booking

import "sync"

type ConcStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcStore() *ConcStore {
	return &ConcStore{
		bookings: map[string]Booking{},
	}
}

func (m *ConcStore) Book(b Booking) error {
	m.Lock()
	defer m.Unlock()

	if _, exists := m.bookings[b.SeatId]; exists {
		return ErrSeatTaken
	}

	m.bookings[b.SeatId] = b
	return nil
}

func (m *ConcStore) ListBookings(MovieId string) []Booking {
	m.RLock()
	defer m.RUnlock()

	var res []Booking

	for _, b := range m.bookings {
		if b.MovieId == MovieId {
			res = append(res, b)
		}
	}
	return res
}
