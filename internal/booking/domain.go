package booking

import (
	"time"
)

// var(
// 	ErrSeatTaken = errors.New("Seat is already taken")
// )

// this is the actual booking struct
type Booking struct {
	ID        string
	MovieId   string
	SeatId    string
	UserId    string
	Status    string
	ExpiresAt time.Time
}

// this is the interface that a struct needs to implement
type BookingInterface interface {
	Book(b Booking) error
	ListBookings(MovieId string) []Booking
}
