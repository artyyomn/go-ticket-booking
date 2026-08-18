package booking

import (
	"errors"
)

var (
	ErrSeatTaken = errors.New("Seat already taken")
)

type MemStore struct {
	//maps seat to booking????
	bookings map[string]Booking
}

func NewMemStore() *MemStore {
	return &MemStore{
		bookings: map[string]Booking{},
	}
}

//implementation of the structs
func (m *MemStore) Book(b Booking) error {
	if _, exists := m.bookings[b.SeatId]; exists {
		return ErrSeatTaken
	}

	m.bookings[b.SeatId] = b
	return nil
}

func (m *MemStore) ListBookings(MovieId string) []Booking {
	var res []Booking

	for _, b := range m.bookings {
		if b.MovieId == MovieId {
			res = append(res, b)
		}
	}
	return res
}
