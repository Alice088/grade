package booking

type BookingService struct {
	Repo BookingRepository
}

func NewService(repo BookingRepository) BookingService {
	return BookingService{
		Repo: repo,
	}
}
