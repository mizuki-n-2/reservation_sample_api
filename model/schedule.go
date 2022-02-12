package model

import "time"

type Schedule struct {
	ID           string `json:"id" gorm:"primaryKey;size:36"`
	Date         string `json:"date" gorm:"not null;size:10"`
	StartTime    string `json:"start_time" gorm:"not null;size:5"`
	MaxNumber    int    `json:"max_number" gorm:"not null"`
	Reservations []Reservation
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
