package model

import (
	"time"
)

type Schedule struct {
	ID                string `json:"id"`
	Date              string `json:"date"`
	StartTime         string `json:"start_time"`
	MaxNumber         int    `json:"max_number"`
	// いらないかも？
	ReservationNumber int    `json:"reservation_number"`
	Reservations      []Reservation
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
