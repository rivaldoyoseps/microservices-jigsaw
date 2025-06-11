package repository

import (
	"github.com/rivaldoyoseps/schedule-service/internal/domain"
	"gorm.io/gorm"
)

type scheduleRepositoryImpl struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepositoryImpl{db}
}

func (r *scheduleRepositoryImpl) Create(slot *domain.Slot) error {
	return r.db.Create(slot).Error
}

func (r *scheduleRepositoryImpl) FindByUserID(userID string) ([]domain.Slot, error) {
	var slots []domain.Slot
	err := r.db.Where("user_id = ?", userID).Find(&slots).Error
	return slots, err
}
