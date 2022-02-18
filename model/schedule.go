package model

import (
	"github.com/google/uuid"
	"time"
)

type Schedule struct {
	ID                string        `json:"id"`
	Date              string        `json:"date"`
	StartTime         string        `json:"start_time"`
	MaxNumber         int           `json:"max_number"`
	Reservations      []Reservation `gorm:"foreignKey:ScheduleID"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

func NewSchedule(date, startTime string, maxNumber int) (*Schedule, error) {
	// TODO: 作成時の(引数の)バリデーション
	schedule := &Schedule{
		ID:        uuid.NewString(),
		Date:      date,
		StartTime: startTime,
		MaxNumber: maxNumber,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return schedule, nil
}
