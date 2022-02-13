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
	ReservationNumber int    `json:"reservation_number"`
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

func GetAvailableSchedules() ([]Schedule, error) {
	var schedules []Schedule

	query := "SELECT * FROM schedules LEFT OUTER JOIN (SELECT schedule_id, COUNT(*) AS reservation_number FROM reservations GROUP BY schedule_id) ON id = schedule_id"

	if err := db.Raw(query).Scan(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

func GetAvailableSchedule(id string) (Schedule, error) {
	var schedule Schedule

	query := "SELECT * FROM schedules LEFT OUTER JOIN (SELECT schedule_id, COUNT(*) AS reservation_number FROM reservations GROUP BY schedule_id) ON id = schedule_id WHERE id = ?"

	if err := db.Raw(query, id).Scan(&schedule).Error; err != nil {
		return Schedule{}, err
	}

	return schedule, nil
}

func UpdateAvailableScheduleMaxNumber(id string, maxNumber int) error {
	db := GetDB()

	if err := db.Model(&Schedule{}).Where("id = ?", id).Update("max_number", maxNumber).Error; err != nil {
		return err
	}
	
	return nil
}

func DeleteAvailableSchedule(id string) error {
	db := GetDB()

	if err := db.Model(&Schedule{}).Where("id = ?", id).Delete(&Schedule{}).Error; err != nil {
		return err
	}

	return nil
}