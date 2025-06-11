package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rivaldoyoseps/schedule-service/internal/domain"
	"github.com/rivaldoyoseps/schedule-service/internal/usecase"
)

type ScheduleHTTPHandler struct {
	usecase usecase.ScheduleUseCase
}

func NewScheduleHTTPHandler(router fiber.Router, uc usecase.ScheduleUseCase) {
	handler := &ScheduleHTTPHandler{usecase: uc}

	router.Post("/slots", handler.CreateSlot)
	router.Get("/slots/user/:userId", handler.GetSlotsByUser)
}

type createSlotRequest struct {	
	UserID    string `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (h *ScheduleHTTPHandler) CreateSlot(c *fiber.Ctx) error {
	var req createSlotRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid start_time format"})
	}
	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid end_time format"})
	}

	slot := &domain.Slot{
		UserID:    req.UserID,
		StartTime: startTime,
		EndTime:   endTime,
		IsBooked:  false,
	}
	if err := h.usecase.CreateSlot(slot); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "slot created"})
}

func (h *ScheduleHTTPHandler) GetSlotsByUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	slots, err := h.usecase.GetSlotsByUser(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(slots)
}
