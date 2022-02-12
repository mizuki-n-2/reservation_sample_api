package model

import "time"

type Reservation struct {
	ID                       string `json:"id" gorm:"primaryKey;size:36"`
	Name                     string `json:"name" gorm:"not null;size:20"`
	Email                    string `json:"email" gorm:"not null;unique"`
	PhoneNumber              string `json:"phone_number" gorm:"not null;size:15"`
	Address                  string `json:"address" gorm:"not null;size:30"`
	AdultNumber              int    `json:"adult_number" gorm:"not null;default:0"`
	PrimarySchoolChildNumber int    `json:"primary_school_child_number" gorm:"not null;default:0"`
	ChildNumber              int    `json:"child_number" gorm:"not null;default:0"`
	SearchID                 string `json:"search_id" gorm:"not null;size:36"`
	ScheduleID               string `json:"schedule_id" gorm:"not null;size:36"`
	Schedule                 Schedule
	CreatedAt                time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt                time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
