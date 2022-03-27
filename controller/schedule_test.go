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
	"github.com/mizuki-n-2/reservation_sample_api/service"
	"github.com/stretchr/testify/assert"
)

func TestSchedule_GetSchedules(t *testing.T) {
	var (
		schedules = []model.Schedule{
			{
				ID:        "schedule-id-1",
				Date:      "2022-01-01",
				StartTime: "10:00",
				Reservations: []model.Reservation{
					{
						ID:         "reservation-id-1",
						ScheduleID: "schedule-id-1",
					},
				},
				MaxNumber: 10,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:           "schedule-id-2",
				Date:         "2022-12-01",
				StartTime:    "09:00",
				Reservations: []model.Reservation{},
				MaxNumber:    15,
				CreatedAt:    time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
			},
		}
		expected = []controller.ScheduleResponse{
			{
				ID:                "schedule-id-1",
				Date:              "2022-01-01",
				StartTime:         "10:00",
				ReservationNumber: 1,
				MaxNumber:         10,
				CreatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:                "schedule-id-2",
				Date:              "2022-12-01",
				StartTime:         "09:00",
				ReservationNumber: 0,
				MaxNumber:         15,
				CreatedAt:         time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:         time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/schedules", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().FindAll().Return(schedules, nil)

	c := controller.NewScheduleController(m1, m2)
	if err := c.GetSchedules()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var actual []controller.ScheduleResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}

func TestSchedule_GetSchedule(t *testing.T) {
	var (
		schedule = model.Schedule{
			ID:        "schedule-id-1",
			Date:      "2022-01-01",
			StartTime: "10:00",
			Reservations: []model.Reservation{
				{
					ID:         "reservation-id-1",
					ScheduleID: "schedule-id-1",
				},
			},
			MaxNumber: 10,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		expected = controller.ScheduleResponse{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			ReservationNumber: 1,
			MaxNumber:         10,
			CreatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/schedules/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(schedule.ID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().FindByID(schedule.ID).Return(schedule, nil)

	c := controller.NewScheduleController(m1, m2)
	if err := c.GetSchedule()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var actual controller.ScheduleResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}

func TestSchedule_CreateSchedule(t *testing.T) {
	var (
		request = controller.ScheduleRequest{
			Date:         "2022-01-01",
			StartTime:    "10:00",
			MaxNumber:    10,
		}
		schedule = model.Schedule{
			ID:                "created-schedule-id",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         10,
		}
		expected = controller.ScheduleResponse{
			ID:                "created-schedule-id",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         10,
			ReservationNumber: 0,
		}
	)

	e := echo.New()
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/auth/schedules", strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m1.EXPECT().CheckAuth(context).Return(nil)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().Create(gomock.Any()).Return(schedule, nil)

	c := controller.NewScheduleController(m1, m2)
	if err := c.CreateSchedule()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)

	var actual controller.ScheduleResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}

func TestSchedule_UpdateSchedule(t *testing.T) {
	var (
		requestID = "schedule-id-1"
		request = controller.ScheduleRequest{
			MaxNumber:    15,
		}
		oldSchedule = model.Schedule{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         10,
		}
		updatedSchedule = model.Schedule{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         15,
		}
		expected = controller.ScheduleResponse{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         15,
			ReservationNumber: 0,
		}
	)

	e := echo.New()
	requestJSON, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/auth/schedules/:id", strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(requestID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m1.EXPECT().CheckAuth(context).Return(nil)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().FindByID(requestID).Return(oldSchedule, nil)
	m2.EXPECT().Update(gomock.Any()).Return(updatedSchedule, nil)

	c := controller.NewScheduleController(m1, m2)
	if err := c.UpdateSchedule()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var actual controller.ScheduleResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expected, actual)
}

func TestSchedule_DeleteSchedule(t *testing.T) {
	var (
		requestID = "schedule-id-1"
		schedule = model.Schedule{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			MaxNumber:         10,
		}
	)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/auth/schedules/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(requestID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m1 := service.NewMockAuthService(ctrl)
	m1.EXPECT().CheckAuth(context).Return(nil)
	m2 := repository.NewMockScheduleRepository(ctrl)
	m2.EXPECT().FindByID(requestID).Return(schedule, nil)
	m2.EXPECT().Delete(schedule.ID).Return(nil)

	c := controller.NewScheduleController(m1, m2)
	if err := c.DeleteSchedule()(context); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Equal(t, "null\n", rec.Body.String())
}
