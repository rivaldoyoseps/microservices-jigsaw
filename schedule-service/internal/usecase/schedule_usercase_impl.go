package usecase

import (
	"github.com/rivaldoyoseps/schedule-service/internal/domain"
	"github.com/rivaldoyoseps/schedule-service/internal/repository"
)

type scheduleUseCase struct {
	repo repository.ScheduleRepository
}

func NewScheduleUseCase(repo repository.ScheduleRepository) ScheduleUseCase {
	return &scheduleUseCase{repo}
}

func (uc *scheduleUseCase) CreateSlot(slot *domain.Slot) error {
	return uc.repo.Create(slot)
}

func (uc *scheduleUseCase) GetSlotsByUser(userID string) ([]domain.Slot, error) {
	return uc.repo.FindByUserID(userID)
}
