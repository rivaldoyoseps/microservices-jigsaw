package domain

import "time"

type Booking struct {
	ID string	`json:"id" gorm:"primaryKey"`
	UserID string `json:"user_id"`
	ScheduleId string `json:"schedule_id"`
	Note string `json:"note"`
	Status string `json:"status"` 		
	CreatetAt time.Time `json:"created_at"`
}