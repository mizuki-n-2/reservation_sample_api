package model_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"golang.org/x/crypto/bcrypt"
)

func TestAdmin_NewAdmin(t *testing.T) {
	type args struct {
		name     string
		email    string
		password string
	}

	tests := []struct {
		name      string
		args      args
		wantAdmin *model.Admin
		wantErr   error
	}{
		{
			name: "正常系",
			args: args{
				name:     "マイケル・ジョン",
				email:    "user@example.com",
				password: "password123",
			},
			wantAdmin: &model.Admin{
				ID:        uuid.NewString(),
				Name:      "マイケル・ジョン",
				Email:     "user@example.com",
				Password:  "password123",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdmin, err := model.NewAdmin(tt.args.name, tt.args.email, tt.args.password)
			if err != nil {
				t.Errorf("NewAdmin() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO: もう少しいい方法を考える
			_, ok := reflect.TypeOf(*gotAdmin).FieldByName("ID")
			if !ok || reflect.TypeOf(gotAdmin.ID) != reflect.TypeOf(tt.wantAdmin.ID) {
				t.Errorf("NewAdmin() ID field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotAdmin).FieldByName("CreatedAt")
			if !ok || reflect.TypeOf(gotAdmin.CreatedAt) != reflect.TypeOf(tt.wantAdmin.CreatedAt) {
				t.Errorf("NewAdmin() CreatedAt field not found or type is not match")
			}
			_, ok = reflect.TypeOf(*gotAdmin).FieldByName("UpdatedAt")
			if !ok || reflect.TypeOf(gotAdmin.UpdatedAt) != reflect.TypeOf(tt.wantAdmin.UpdatedAt) {
				t.Errorf("NewAdmin() UpdatedAt field not found or type is not match")
			}
			if gotAdmin.Name != tt.wantAdmin.Name {
				t.Errorf("NewAdmin() gotAdmin.Name = %v, want %v", gotAdmin.Name, tt.wantAdmin.Name)
			}
			if gotAdmin.Email != tt.wantAdmin.Email {
				t.Errorf("NewAdmin() gotAdmin.Email = %v, want %v", gotAdmin.Email, tt.wantAdmin.Email)
			}
			if err = bcrypt.CompareHashAndPassword([]byte(gotAdmin.Password), []byte(tt.wantAdmin.Password)); err != nil {
					t.Errorf("NewAdmin() gotAdmin.Password = %v, want %v", gotAdmin.Password, tt.wantAdmin.Password)
				}
		})
	}
}

func TestAdmin_NewName(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name     string
		args     args
		wantName model.Name
		wantErr  error
	}{
		{
			name: "正常系",
			args: args{
				value: "マイケル・ジョン",
			},
			wantName: model.Name("マイケル・ジョン"),
			wantErr:  nil,
		},
		{
			name: "異常系：文字数が短い場合エラーを返す",
			args: args{
				value: "a",
			},
			wantName: "",
			wantErr:  fmt.Errorf("nameは%d文字以上%d文字以下にしてください", model.MIN_LENGTH_USER_NAME, model.MAX_LENGTH_USER_NAME),
		},
		{
			name: "異常系：文字数が長い場合エラーを返す",
			args: args{
				value: "マイケル・マイケルマイケル・マイケルマイケル",
			},
			wantName: "",
			wantErr:  fmt.Errorf("nameは%d文字以上%d文字以下にしてください", model.MIN_LENGTH_USER_NAME, model.MAX_LENGTH_USER_NAME),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, err := model.NewName(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotName != tt.wantName {
				t.Errorf("NewName() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestAdmin_NewEmail(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name      string
		args      args
		wantEmail model.Email
		wantErr   error
	}{
		{
			name: "正常系",
			args: args{
				value: "user@example.com",
			},
			wantEmail: model.Email("user@example.com"),
			wantErr:   nil,
		},
		{
			name: "異常系：@がない場合エラーを返す",
			args: args{
				value: "userexample.com",
			},
			wantEmail: "",
			wantErr:   fmt.Errorf("emailの形式が正しくありません"),
		},
		{
			name: "異常系：@の前がない場合エラーを返す",
			args: args{
				value: "@example.com",
			},
			wantEmail: "",
			wantErr:   fmt.Errorf("emailの形式が正しくありません"),
		},
		{
			name: "異常系：@の後がない場合エラーを返す",
			args: args{
				value: "user@",
			},
			wantEmail: "",
			wantErr:   fmt.Errorf("emailの形式が正しくありません"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEmail, err := model.NewEmail(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotEmail != tt.wantEmail {
				t.Errorf("NewEmail() gotEmail = %v, want %v", gotEmail, tt.wantEmail)
			}
		})
	}
}

func TestAdmin_NewPassword(t *testing.T) {
	type args struct {
		value string
	}

	tests := []struct {
		name         string
		args         args
		wantPassword model.Password
		wantErr      error
	}{
		{
			name: "正常系",
			args: args{
				value: "password123",
			},
			wantPassword: "password123",
			wantErr:      nil,
		},
		{
			name: "異常系：文字数が短い場合エラーを返す",
			args: args{
				value: "pass",
			},
			wantPassword: "",
			wantErr:      fmt.Errorf("passwordは%d文字以上%d文字以下にしてください", model.MIN_LENGTH_USER_PASSWORD, model.MAX_LENGTH_USER_PASSWORD),
		},
		{
			name: "異常系：文字数が長い場合エラーを返す",
			args: args{
				value: "password1234567password12345678",
			},
			wantPassword: "",
			wantErr:      fmt.Errorf("passwordは%d文字以上%d文字以下にしてください", model.MIN_LENGTH_USER_PASSWORD, model.MAX_LENGTH_USER_PASSWORD),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPassword, err := model.NewPassword(tt.args.value)
			// TODO: エラーを文字列比較でない方法にする
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("NewPassword() error = %v, wantErr %v", err, tt.wantErr)
			}

			if gotPassword != "" {
				if err = bcrypt.CompareHashAndPassword([]byte(gotPassword), []byte(tt.wantPassword)); err != nil {
					t.Errorf("NewPassword() gotPassword = %v, want %v", gotPassword, tt.wantPassword)
				}
			}
		})
	}
}
