package repository

import (
	"errors"

	"github.com/rivaldoyoseps/booking-service/internal/domain"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking *domain.Booking) error
	FindByUser(userID string)([]domain.Booking, error)
	FindById(id string) (*domain.Booking, error)
}

type bookingRepository struct{
	DB *gorm.DB
}

func NewBookingRepository (DB *gorm.DB) BookingRepository {
	return &bookingRepository{
		DB: DB,
	}
}

func (repo *bookingRepository) Create(booking *domain.Booking) error {
	return repo.DB.Create(booking).Error
}

func (repo *bookingRepository) FindByUser(userID string)([]domain.Booking, error){
	var bookings []domain.Booking
	err := repo.DB.Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (repo *bookingRepository) FindById(id string) (*domain.Booking, error) {
	var booking domain.Booking
	err := repo.DB.Where("id = ?", id).Take(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &booking, err
}