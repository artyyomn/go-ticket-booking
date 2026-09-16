package booking

import (
	"github.com/artyyomn/go-ticket-booking/internal/utils"
	"net/http"
)

type handler struct {
	svc *Service
}

func NewHandler(svc *Service) *handler {
	return &handler{svc}
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("movieId")
	bookings := h.svc.ListBookings(movieId)

	seats := make([]seatInfo, len(bookings))
	for _, b := range bookings {
		seats = append(seats, seatInfo{
			SeatId: b.SeatId,
			UserId: b.UserId,
			Booked: true,
		})
	}

	utils.WriteJSON(w, r,http.StatusOK, seats)
}

type seatInfo struct {
	SeatId string `json:"seat_id"`
	UserId string `json:"user_id"`
	Booked bool   `json:"booked"`
}
