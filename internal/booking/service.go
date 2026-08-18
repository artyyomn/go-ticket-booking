package booking

type Service struct {
	store BookingInterface
}

func NewService(store BookingInterface) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Book(b Booking) error {
	return s.store.Book(b)
}

func (s *Service) ListBookings(MovieId string) []Booking{
	return s.store.ListBookings(MovieId)
}
