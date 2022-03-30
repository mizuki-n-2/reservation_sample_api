package model

import (
	"github.com/google/uuid"
	"time"
	"fmt"
	"regexp"
)

type Schedule struct {
	ID           string        `json:"id"`
	Date         Date          `json:"date"`
	StartTime    StartTime     `json:"start_time"`
	MaxNumber    MaxNumber     `json:"max_number"`
	Reservations []Reservation `json:"reservations"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func NewSchedule(date, startTime string, maxNumber int) (*Schedule, error) {
	newDate, err := NewDate(date)
	if err != nil {
		return nil, err
	}

	newStartTime, err := NewStartTime(startTime)
	if err != nil {
		return nil, err
	}

	newMaxNumber, err := NewMaxNumber(maxNumber)
	if err != nil {
		return nil, err
	}

	schedule := &Schedule{
		ID:        uuid.NewString(),
		Date:      newDate,
		StartTime: newStartTime,
		MaxNumber: newMaxNumber,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return schedule, nil
}

type Date string

var (
	DATE_PATTERN = `^\d{4}-\d{2}-\d{2}$`
)

func NewDate(value string) (Date, error) {
	if !regexp.MustCompile(DATE_PATTERN).MatchString(value) {
		return "", fmt.Errorf("dateの形式が正しくありません")
	}

	return Date(value), nil
}

type StartTime string

var (
	START_TIME_PATTERN = `^\d{2}:\d{2}$`
)

func NewStartTime(value string) (StartTime, error) {
	if !regexp.MustCompile(START_TIME_PATTERN).MatchString(value) {
		return "", fmt.Errorf("start_timeの形式が正しくありません")
	}

	return StartTime(value), nil
}

type MaxNumber int

var (
	MAX_MAX_NUMBER = 100
	MIN_MAX_NUMBER = 1
)

func NewMaxNumber(value int) (MaxNumber, error) {
	if value < MIN_MAX_NUMBER || value > MAX_MAX_NUMBER {
		return 0, fmt.Errorf("max_numberは%d以上%d以下にしてください", MIN_MAX_NUMBER, MAX_MAX_NUMBER)
	}

	return MaxNumber(value), nil
}

func (schedule *Schedule) UpdateMaxNumber(maxNumber int) error {
	newMaxNumber, err := NewMaxNumber(maxNumber)
	if err != nil {
		return err
	}

	schedule.MaxNumber = newMaxNumber
	schedule.UpdatedAt = time.Now()

	return nil
}
