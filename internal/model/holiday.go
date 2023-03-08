package model

import (
	"time"
)

type Holiday struct {
	Id       uint16    `json:"id"`
	Type     string    `json:"type"`
	UserId   string    `json:"userId"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	FromDate time.Time `json:"fromDate"`
	ToDate   time.Time `json:"toDate"`
	Creation time.Time `json:"date"`
}

type HolidayRepository interface {
	FindAll() ([]Holiday, error)
	Save(holiday Holiday) (Holiday, error)
	FindById(id uint) (Holiday, error)
	DeleteHoliday(id uint) (int64, error)
}
