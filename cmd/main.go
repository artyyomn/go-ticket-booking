package main

import (
	"log"
	"net/http"

	"github.com/artyyomn/go-ticket-booking/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/artyyomn/go-ticket-booking/internal/booking"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", ListMovies)
	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	svc := booking.NewService(store)
	bookingHanlder := booking.NewHandler(svc)

	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHanlder.ListSeats)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error running the server", err)
	}
}

var movies = []MovieResponse{
	{ID: "DUNE 3", Title: "DUNE 3", Rows: 15, SeatsPerRow: 10},
	{ID: "DOOMSDAY", Title: "DOOMSDAY", Rows: 15, SeatsPerRow: 10},
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, r, http.StatusOK, movies)
}

type MovieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}
