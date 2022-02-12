package model

import "time"

type Admin struct {
	ID        string    `json:"id" gorm:"primaryKey;size:36"`
	Name      string    `json:"name" gorm:"not null;size:20"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null;size:30"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}
