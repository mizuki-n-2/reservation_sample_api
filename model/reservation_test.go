package model_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/mizuki-n-2/reservation_sample_api/model"
)

func TestReservation_NewReservation(t *testing.T) {
	type args struct {
		name                     string
		email                    string
		phoneNumber              string
		address                  string
		adultNumber              int
		primarySchoolChildNumber int
		childNumber              int
		scheduleID               string
	}

	tests := []struct {
		name            string
		args            args
		wantReservation *model.Reservation
		wantErr         error
	}{
		{
			name: "正常系",
			args: args{
				name:     "山田太郎",
				email:    "user@example.com",
				phoneNumber: "090-1234-5678",
				address: "東京都",
				adultNumber: 3,
				primarySchoolChildNumber: 2,
				childNumber: 0,
				scheduleID: "b17d9dcc-5149-41e6-b427-52e0049b7f9e",
			},
			wantReservation: &model.Reservation{
				ID:        uuid.NewString(),
				Name:      "山田太郎",
				Email:     "user@example.com",
				PhoneNumber: "090-1234-5678",
				Address: "東京都",
				AdultNumber: 3,
				PrimarySchoolChildNumber: 2,
				ChildNumber: 0,
				SearchID: uuid.NewString(),
				ScheduleID: "b17d9dcc-5149-41e6-b427-52e0049b7f9e",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReservation, err := model.NewReservation(tt.args.name, tt.args.email, tt.args.phoneNumber, tt.args.address, tt.args.adultNumber, tt.args.primarySchoolChildNumber, tt.args.childNumber, tt.args.scheduleID)
			if err != nil {
				t.Errorf("NewReservation() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO: もう少しいい方法を考える
			_, ok := reflect.TypeOf(*gotReservation).FieldByName("ID")
			if !ok || reflect.TypeOf(gotReservation.ID) != reflect.TypeOf(tt.wantReservation.ID) {
				t.Errorf("NewReservation() ID field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotReservation).FieldByName("CreatedAt")
			if !ok || reflect.TypeOf(gotReservation.CreatedAt) != reflect.TypeOf(tt.wantReservation.CreatedAt) {
				t.Errorf("NewReservation() CreatedAt field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotReservation).FieldByName("UpdatedAt")
			if !ok || reflect.TypeOf(gotReservation.UpdatedAt) != reflect.TypeOf(tt.wantReservation.UpdatedAt) {
				t.Errorf("NewReservation() UpdatedAt field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotReservation).FieldByName("SearchID")
			if !ok || reflect.TypeOf(gotReservation.SearchID) != reflect.TypeOf(tt.wantReservation.SearchID) {
				t.Errorf("NewReservation() SearchID field not found or type is not match")
			}
			if gotReservation.Name != tt.wantReservation.Name {
				t.Errorf("NewReservation() gotReservation.Name = %v, want %v", gotReservation.Name, tt.wantReservation.Name)
			}
			if gotReservation.Email != tt.wantReservation.Email {
				t.Errorf("NewReservation() gotReservation.Email = %v, want %v", gotReservation.Email, tt.wantReservation.Email)
			}
			if gotReservation.PhoneNumber != tt.wantReservation.PhoneNumber {
				t.Errorf("NewReservation() gotReservation.PhoneNumber = %v, want %v", gotReservation.PhoneNumber, tt.wantReservation.PhoneNumber)
			}
			if gotReservation.Address != tt.wantReservation.Address {
				t.Errorf("NewReservation() gotReservation.Address = %v, want %v", gotReservation.Address, tt.wantReservation.Address)
			}
			if gotReservation.AdultNumber != tt.wantReservation.AdultNumber {
				t.Errorf("NewReservation() gotReservation.AdultNumber = %v, want %v", gotReservation.AdultNumber, tt.wantReservation.AdultNumber)
			}
			if gotReservation.PrimarySchoolChildNumber != tt.wantReservation.PrimarySchoolChildNumber {
				t.Errorf("NewReservation() gotReservation.PrimarySchoolChildNumber = %v, want %v", gotReservation.PrimarySchoolChildNumber, tt.wantReservation.PrimarySchoolChildNumber)
			}
			if gotReservation.ChildNumber != tt.wantReservation.ChildNumber {
				t.Errorf("NewReservation() gotReservation.ChildNumber = %v, want %v", gotReservation.ChildNumber, tt.wantReservation.ChildNumber)
			}
			if gotReservation.ScheduleID != tt.wantReservation.ScheduleID {
				t.Errorf("NewReservation() gotReservation.ScheduleID = %v, want %v", gotReservation.ScheduleID, tt.wantReservation.ScheduleID)
			}
		})
	}
}
func TestReservation_NewPhoneNumber(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name            string
		args            args
		wantPhoneNumber model.PhoneNumber
		wantErr         error
	}{
		{
			name: "正常系",
			args: args{
				value: "090-1234-5678",
			},
			wantPhoneNumber: model.PhoneNumber("090-1234-5678"),
			wantErr:         nil,
		},
		{
			name: "異常系：ハイフンなしの場合はエラーを返す",
			args: args{
				value: "09012345678",
			},
			wantPhoneNumber: "",
			wantErr:         fmt.Errorf("電話番号の形式が正しくありません"),
		},
		{
			name: "異常系：0から始まらない場合はエラーを返す",
			args: args{
				value: "99012345678",
			},
			wantPhoneNumber: "",
			wantErr:         fmt.Errorf("電話番号の形式が正しくありません"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPhoneNumber, err := model.NewPhoneNumber(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotPhoneNumber != tt.wantPhoneNumber {
				t.Errorf("NewPhoneNumber() gotPhoneNumber = %v, want %v", gotPhoneNumber, tt.wantPhoneNumber)
			}
		})
	}
}
func TestReservation_NewAddress(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name        string
		args        args
		wantAddress model.Address
		wantErr     error
	}{
		{
			name: "正常系",
			args: args{
				value: "東京都",
			},
			wantAddress: model.Address("東京都"),
			wantErr:     nil,
		},
		{
			name: "異常系：住所がない場合はエラーを返す",
			args: args{
				value: "",
			},
			wantAddress: "",
			wantErr:     fmt.Errorf("住所が入力されていません"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, err := model.NewAddress(tt.args.value)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("NewAddress() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
		})
	}
}
func TestReservation_NewNumberOfPeople(t *testing.T) {
	type args struct {
		value int
	}

	tests := []struct {
		name               string
		args               args
		wantNumberOfPeople model.NumberOfPeople
		wantErr            error
	}{
		{
			name: "正常系",
			args: args{
				value: 10,
			},
			wantNumberOfPeople: model.NumberOfPeople(10),
			wantErr:            nil,
		},
		{
			name: "異常系：負の数字の場合はエラーを返す",
			args: args{
				value: -2,
			},
			wantNumberOfPeople: -1,
			wantErr:            fmt.Errorf("人数は%d人以上%d人以下にしてください", model.MIN_NUMBER_OF_PEOPLE, model.MAX_NUMBER_OF_PEOPLE),
		},
		{
			name: "異常系：数字が大きい場合はエラーを返す",
			args: args{
				value: 100,
			},
			wantNumberOfPeople: -1,
			wantErr:            fmt.Errorf("人数は%d人以上%d人以下にしてください", model.MIN_NUMBER_OF_PEOPLE, model.MAX_NUMBER_OF_PEOPLE),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumberOfPeople, err := model.NewNumberOfPeople(tt.args.value)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewNumberOfPeople() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotNumberOfPeople != tt.wantNumberOfPeople {
				t.Errorf("NewNumberOfPeople() gotNumberOfPeople = %v, want %v", gotNumberOfPeople, tt.wantNumberOfPeople)
			}
		})
	}
}
