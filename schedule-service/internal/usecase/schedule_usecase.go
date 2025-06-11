package usecase

import "github.com/rivaldoyoseps/schedule-service/internal/domain"

type ScheduleUseCase interface {
	CreateSlot(slot *domain.Slot) error
	GetSlotsByUser(userID string) ([]domain.Slot, error)
}
