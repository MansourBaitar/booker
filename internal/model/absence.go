package model

import (
	"time"
)

type Absence struct {
	Id       uint16    `json:"id"`
	Type     string    `json:"type"`
	UserId   string    `json:"userId"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	FromDate time.Time `json:"fromDate"`
	ToDate   time.Time `json:"toDate"`
	Creation time.Time `json:"date"`
}

type AbsenceRepository interface {
	FindAll() ([]Absence, error)
	Save(expense Absence) (Absence, error)
	FindById(id uint) (Absence, error)
	DeleteAbsence(id uint) (int64, error)
}
