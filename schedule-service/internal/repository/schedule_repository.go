package repository

import "github.com/rivaldoyoseps/schedule-service/internal/domain"

type ScheduleRepository interface {
	Create(slot *domain.Slot) error
	FindByUserID(userID string) ([]domain.Slot, error)
}
