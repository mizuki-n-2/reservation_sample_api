package model

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Reservation struct {
	ID                       string    `json:"id" gorm:"primaryKey;size:36"`
	Name                     string    `json:"name" gorm:"not null;size:20"`
	Email                    string    `json:"email" gorm:"not null;unique"`
	PhoneNumber              string    `json:"phone_number" gorm:"not null;size:15"`
	Address                  string    `json:"address" gorm:"not null;size:30"`
	AdultNumber              int       `json:"adult_number" gorm:"not null;default:0"`
	PrimarySchoolChildNumber int       `json:"primary_school_child_number" gorm:"not null;default:0"`
	ChildNumber              int       `json:"child_number" gorm:"not null;default:0"`
	SearchID                 string    `json:"search_id" gorm:"not null;size:36"`
	ScheduleID               string    `json:"schedule_id" gorm:"not null;size:36"`
	Date                     string    `json:"date"`
	StartTime                string    `json:"start_time"`
	CreatedAt                time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt                time.Time `json:"updated_at" gorm:"autoUpdateTime;not null"`
}

func GetReservations() ([]Reservation, error) {
	db := GetDB()

	var reservations []Reservation

	query := "SELECT * FROM reservations JOIN schedules ON reservations.schedule_id = schedules.id"

	if err := db.Raw(query).Scan(&reservations).Error; err != nil {
		return nil, err
	}

	log.Println(reservations)

	return reservations, nil
}

func CreateReservation(name, email, phoneNumber, address string, adultNumber, primarySchoolChildNumber, childNumber int, scheduleID string) (string, error) {
	db := GetDB()

	reservation := Reservation{
		ID:                       uuid.NewString(),
		Name:                     name,
		Email:                    email,
		PhoneNumber:              phoneNumber,
		Address:                  address,
		AdultNumber:              adultNumber,
		PrimarySchoolChildNumber: primarySchoolChildNumber,
		ChildNumber:              childNumber,
		SearchID:                 uuid.NewString(),
		ScheduleID:               scheduleID,
	}

	if err := db.Create(&reservation).Error; err != nil {
		return "", err
	}

	return reservation.ID, nil
}


func GetReservation(id string) (Reservation, error) {
	db := GetDB()

	var reservation Reservation

	query := "SELECT * FROM reservations JOIN schedules ON reservations.schedule_id = schedules.id WHERE reservations.id = ?"

	if err := db.Raw(query, id).Scan(&reservation).Error; err != nil {
		return Reservation{}, err
	}

	return reservation, nil
}

func DeleteReservation(id string) error {
	db := GetDB()

	if err := db.Where("id = ?", id).Delete(&Reservation{}).Error; err != nil {
		return err
	}

	return nil
}
