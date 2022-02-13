package model

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID                string `json:"id" gorm:"primaryKey;size:36"`
	Date              string `json:"date" gorm:"not null;size:10"`
	StartTime         string `json:"start_time" gorm:"not null;size:5"`
	MaxNumber         int    `json:"max_number" gorm:"not null;default:0"`
	ReservationNumber int    `json:"reservation_number" gorm:"-"`
	Reservations      []Reservation
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}

func CreateAvailableSchedule(date, startTime string, maxNumber int) string {
	db := GetDB()
	
	schedule := Schedule{
		ID: uuid.NewString(),
		Date: date,
		StartTime: startTime,
		MaxNumber: maxNumber,
	}

	db.Create(&schedule)

	return schedule.ID
}