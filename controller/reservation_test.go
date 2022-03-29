package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/mizuki-n-2/reservation_sample_api/controller"
	"github.com/mizuki-n-2/reservation_sample_api/model"
	"github.com/mizuki-n-2/reservation_sample_api/repository"
	v "github.com/mizuki-n-2/reservation_sample_api/validator"
	"github.com/stretchr/testify/assert"
	"github.com/go-playground/validator/v10"
)

func TestReservation_GetReservations(t *testing.T) {
	var (
		reservations = []model.Reservation{
			{
				ID:                       "reservation-id-1",
				Name:                     "user1",
				Email:                    "user1@example.com",
				PhoneNumber:              "090-1234-5678",
				Address:                  "東京都",
				AdultNumber:              2,
				PrimarySchoolChildNumber: 0,
				ChildNumber:              1,
				SearchID:                 "search-id-1",
				ScheduleID:               "schedule-id-1",
				CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:                       "reservation-id-2",
				Name:                     "user2",
				Email:                    "user2@example.com",
				PhoneNumber:              "090-1234-5678",
				Address:                  "東京都",
				AdultNumber:              1,
				PrimarySchoolChildNumber: 2,
				ChildNumber:              0,
				SearchID:                 "search-id-2",
				ScheduleID:               "schedule-id-1",
				CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reservations", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := repository.NewMockReservationRepository(ctrl)
	m1.EXPECT().FindAll().Return(reservations, nil)
	m2 := repository.NewMockScheduleRepository(ctrl)

	c := controller.NewReservationController(m1, m2)
	if err := c.GetReservations()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var actual []model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, reservations, actual)
}

func TestReservation_GetReservation(t *testing.T) {
	var (
		reservation = model.Reservation{
			ID:                       "reservation-id-1",
			Name:                     "user1",
			Email:                    "user1@example.com",
			PhoneNumber:              "090-1234-5678",
			Address:                  "東京都",
			AdultNumber:              2,
			PrimarySchoolChildNumber: 0,
			ChildNumber:              1,
			SearchID:                 "search-id-1",
			ScheduleID:               "schedule-id-1",
			CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reservations/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(reservation.ID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := repository.NewMockReservationRepository(ctrl)
	m1.EXPECT().FindByID(reservation.ID).Return(reservation, nil)
	m2 := repository.NewMockScheduleRepository(ctrl)

	c := controller.NewReservationController(m1, m2)
	if err := c.GetReservation()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var actual model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, reservation, actual)
}

func TestReservation_DeleteReservation(t *testing.T) {
	var (
		reservations = []model.Reservation{
			{
				ID:                       "reservation-id-1",
				Name:                     "user1",
				Email:                    "user1@example.com",
				PhoneNumber:              "090-1234-5678",
				Address:                  "東京都",
				AdultNumber:              2,
				PrimarySchoolChildNumber: 0,
				ChildNumber:              1,
				SearchID:                 "search-id-1",
				ScheduleID:               "schedule-id-1",
				CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:                       "reservation-id-2",
				Name:                     "user2",
				Email:                    "user2@example.com",
				PhoneNumber:              "090-1234-5678",
				Address:                  "東京都",
				AdultNumber:              1,
				PrimarySchoolChildNumber: 2,
				ChildNumber:              0,
				SearchID:                 "search-id-2",
				ScheduleID:               "schedule-id-1",
				CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/reservations/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(reservations[0].ID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := repository.NewMockReservationRepository(ctrl)
	m1.EXPECT().FindByID(reservations[0].ID).Return(reservations[0], nil)
	m1.EXPECT().Delete(reservations[0].ID).Return(nil)
	m2 := repository.NewMockScheduleRepository(ctrl)

	c := controller.NewReservationController(m1, m2)
	if err := c.DeleteReservation()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "null\n", rec.Body.String())
}

func TestReservation_CreateReservation(t *testing.T) {
	var (
		request = controller.ReservationRequest{
			Name:                     "user1",
			Email:                    "user1@example.com",
			PhoneNumber:              "090-1234-5678",
			Address:                  "東京都",
			AdultNumber:              2,
			PrimarySchoolChildNumber: 0,
			ChildNumber:              1,
			ScheduleID: 						 "38e4cf39-19c3-4a5e-8546-c1dbe1b701dd",
		}
		reservation = model.Reservation{
			ID:                       "created-reservation-id",
			Name:                     "user1",
			Email:                    "user1@example.com",
			PhoneNumber:              "090-1234-5678",
			Address:                  "東京都",
			AdultNumber:              2,
			PrimarySchoolChildNumber: 0,
			ChildNumber:              1,
			ScheduleID: 						 "38e4cf39-19c3-4a5e-8546-c1dbe1b701dd",
			SearchID:                 "created-search-id",
			CreatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		schedule = model.Schedule{
			ID: 									"38e4cf39-19c3-4a5e-8546-c1dbe1b701dd",
		}
	)

	e := echo.New()
	validate := validator.New()
	if err := validate.RegisterValidation("phone", v.PhoneValidator); err != nil {
		t.Fatal(err)
	}
	e.Validator = &v.CustomValidator{Validator: validate}
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/reservations", strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := repository.NewMockReservationRepository(ctrl)
	m1.EXPECT().Create(gomock.Any()).Return(reservation, nil)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().FindByID(request.ScheduleID).Return(schedule, nil)

	c := controller.NewReservationController(m1, m2)
	if err := c.CreateReservation()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)

	var actual model.Reservation
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, reservation, actual)
}