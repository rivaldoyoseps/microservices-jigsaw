package usecase

import (
	"github.com/rivaldoyoseps/booking-service/internal/domain"
	"github.com/rivaldoyoseps/booking-service/internal/repository"
)

type BookingUseCase interface {
	CreateBooking(b *domain.Booking) error
	GetBookingsByUser(userID string) ([]domain.Booking, error)
	GetBookingDetail(id string) (*domain.Booking, error)
}

type bookingUseCase struct {
	repo repository.BookingRepository
}

func NewBookingUseCase (r repository.BookingRepository) BookingUseCase {
	return &bookingUseCase{
		repo: r,
	}
}

func (uc *bookingUseCase) CreateBooking(b *domain.Booking) error {
	return uc.repo.Create(b)
}

func (uc *bookingUseCase) GetBookingsByUser(userID string) ([]domain.Booking, error){
	return uc.repo.FindByUser(userID)
}

func (uc *bookingUseCase) GetBookingDetail(id string) (*domain.Booking, error) {
	return uc.repo.FindById(id)
}

