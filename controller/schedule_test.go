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
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			},
			{
				ID:           "schedule-id-2",
				Date:         "2022-12-01",
				StartTime:    "09:00",
				Reservations: []model.Reservation{},
				MaxNumber:    15,
				CreatedAt:    time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt:    time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local),
			},
		}
		expected = []controller.ScheduleResponse{
			{
				ID:                "schedule-id-1",
				Date:              "2022-01-01",
				StartTime:         "10:00",
				ReservationNumber: 1,
				MaxNumber:         10,
				CreatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			},
			{
				ID:                "schedule-id-2",
				Date:              "2022-12-01",
				StartTime:         "09:00",
				ReservationNumber: 0,
				MaxNumber:         15,
				CreatedAt:         time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local),
				UpdatedAt:         time.Date(2022, 12, 1, 0, 0, 0, 0, time.Local),
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
	m1 := repository.NewMockScheduleRepository(ctrl)
	m1.EXPECT().FindAll().Return(schedules, nil)
	m2 := repository.NewMockAdminRepository(ctrl)

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
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		}
		expected = controller.ScheduleResponse{
			ID:                "schedule-id-1",
			Date:              "2022-01-01",
			StartTime:         "10:00",
			ReservationNumber: 1,
			MaxNumber:         10,
			CreatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
			UpdatedAt:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
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
	m1 := repository.NewMockScheduleRepository(ctrl)
	m1.EXPECT().FindByID(schedule.ID).Return(schedule, nil)
	m2 := repository.NewMockAdminRepository(ctrl)

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
