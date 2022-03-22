package model_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/mizuki-n-2/reservation_sample_api/model"
)

func TestSchedule_NewSchedule(t *testing.T) {
	type args struct {
		date     string
		startTime string
		maxNumber int
	}

	tests := []struct {
		name        string
		args        args
		wantSchedule *model.Schedule
		wantErr      error
	}{
		{
			name: "正常系",
			args: args{
				date:     "2020-01-01",
				startTime: "09:00",
				maxNumber: 10,
			},
			wantSchedule: &model.Schedule{
				Date:     model.Date("2020-01-01"),
				StartTime: model.StartTime("09:00"),
				MaxNumber: model.MaxNumber(10),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSchedule, err := model.NewSchedule(tt.args.date, tt.args.startTime, tt.args.maxNumber)

			if err != nil {
				t.Errorf("NewSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO: もう少しいい方法を考える
			_, ok := reflect.TypeOf(*gotSchedule).FieldByName("ID")
			if !ok || reflect.TypeOf(gotSchedule.ID) != reflect.TypeOf(tt.wantSchedule.ID) {
				t.Errorf("NewSchedule() ID field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotSchedule).FieldByName("CreatedAt")
			if !ok || reflect.TypeOf(gotSchedule.CreatedAt) != reflect.TypeOf(tt.wantSchedule.CreatedAt) {
				t.Errorf("NewSchedule() CreatedAt field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotSchedule).FieldByName("UpdatedAt")
			if !ok || reflect.TypeOf(gotSchedule.UpdatedAt) != reflect.TypeOf(tt.wantSchedule.UpdatedAt) {
				t.Errorf("NewSchedule() UpdatedAt field not found or type is not match")
			}
			if gotSchedule.Date != tt.wantSchedule.Date {
				t.Errorf("NewSchedule() gotSchedule.Date = %v, want %v", gotSchedule.Date, tt.wantSchedule.Date)
			}
			if gotSchedule.StartTime != tt.wantSchedule.StartTime {
				t.Errorf("NewSchedule() gotSchedule.StartTime = %v, want %v", gotSchedule.StartTime, tt.wantSchedule.StartTime)
			}
			if gotSchedule.MaxNumber != tt.wantSchedule.MaxNumber {
				t.Errorf("NewSchedule() gotSchedule.MaxNumber = %v, want %v", gotSchedule.MaxNumber, tt.wantSchedule.MaxNumber)
			}
		})
	}
}

func TestSchedule_NewDate(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name        string
		args        args
		wantDate    model.Date
		wantErr     error
	}{
		{
			name: "正常系",
			args: args{
				value: "2020-01-01",
			},
			wantDate: model.Date("2020-01-01"),
			wantErr:  nil,
		},
		{
			name: "異常系：dateの形式が不正な場合エラーを返す",
			args: args{
				value: "2020-1-1",
			},
			wantDate: "",
			wantErr:  fmt.Errorf("dateの形式が正しくありません"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDate, err := model.NewDate(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewDate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotDate != tt.wantDate {
				t.Errorf("NewDate() gotDate = %v, want %v", gotDate, tt.wantDate)
			}
		})
	}
}

func TestSchedule_NewStartTime(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name        string
		args        args
		wantStartTime model.StartTime
		wantErr      error
	}{
		{
			name: "正常系",
			args: args{
				value: "09:00",
			},
			wantStartTime: model.StartTime("09:00"),
			wantErr: 		 nil,
		},
		{
			name: "異常系：startTimeの形式が不正な場合エラーを返す",
			args: args{
				value: "9:00",
			},
			wantStartTime: "",
			wantErr: 		 fmt.Errorf("start_timeの形式が正しくありません"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartTime, err := model.NewStartTime(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewStartTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotStartTime != tt.wantStartTime {
				t.Errorf("NewStartTime() gotStartTime = %v, want %v", gotStartTime, tt.wantStartTime)
			}
		})
	}
}

func TestSchedule_NewMaxNumber(t *testing.T) {
	type args struct {
		value int
	}

	tests := []struct {
		name        string
		args        args
		wantMaxNumber model.MaxNumber
		wantErr      error
	}{
		{
			name: "正常系",
			args: args{
				value: 10,
			},
			wantMaxNumber: model.MaxNumber(10),
			wantErr: 		 nil,
		},
		{
			name: "異常系：max_numberが0以下の場合エラーを返す",
			args: args{
				value: 0,
			},
			wantMaxNumber: -1,
			wantErr: fmt.Errorf("max_numberは%d以上%d以下にしてください", model.MIN_MAX_NUMBER, model.MAX_MAX_NUMBER),
		},
		{
			name: "異常系：max_numberが大きい場合エラーを返す",
			args: args{
				value: 200,
			},
			wantMaxNumber: -1,
			wantErr: fmt.Errorf("max_numberは%d以上%d以下にしてください", model.MIN_MAX_NUMBER, model.MAX_MAX_NUMBER),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMaxNumber, err := model.NewMaxNumber(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewMaxNumber() error = %v, wantErr %v", err, tt.wantErr)
			}

			if gotMaxNumber != tt.wantMaxNumber {
				t.Errorf("NewMaxNumber() gotMaxNumber = %v, want %v", gotMaxNumber, tt.wantMaxNumber)
			}
		})
	}
}

func TestSchedule_UpdateMaxNumber(t *testing.T) {
	type args struct {
		maxNumber int
	}
	type field struct {
		schedule model.Schedule
	}

	tests := []struct {
		name string
		args args
		field field
		wantErr error
	} {
		{
			name: "正常系",
			args: args{
				maxNumber: 10,
			},
			field: field{
				schedule: model.Schedule{
					ID: "",
					Date: "2022-01-01",
					StartTime: "09:00",
					MaxNumber: 15,
					Reservations: []model.Reservation{},
					CreatedAt: time.Date(2021, 12, 31, 12, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 12, 31, 12, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "異常系：max_numberが不適切な数字の場合エラーを返す",
			args: args{
				maxNumber: 1000,
			},
			field: field{
				schedule: model.Schedule{
					ID: "",
					Date: "2022-01-01",
					StartTime: "09:00",
					MaxNumber: 15,
					Reservations: []model.Reservation{},
					CreatedAt: time.Date(2021, 12, 31, 12, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 12, 31, 12, 0, 0, 0, time.UTC),
				},
			},
			wantErr: fmt.Errorf("max_numberは%d以上%d以下にしてください", model.MIN_MAX_NUMBER, model.MAX_MAX_NUMBER),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := tt.field.schedule
			err := schedule.UpdateMaxNumber(tt.args.maxNumber)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("UpdateMaxNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && schedule.MaxNumber != model.MaxNumber(tt.args.maxNumber) {
				t.Errorf("UpdateMaxNumber() actual = %v, expected = %v", schedule.MaxNumber, tt.args.maxNumber)
			}
		})
	}
}
