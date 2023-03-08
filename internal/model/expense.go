package model

import (
	"time"
)

type Expense struct {
	Id       uint16    `json:"id"`
	Type     string    `json:"type"`
	UserId   string    `json:"userId"`
	Name     string    `json:"name"`
	Amount   float32   `json:"amount"`
	Status   string    `json:"status"`
	Category string    `json:"category"`
	Group    string    `json:"group"`
	Comment  string    `json:"comment"`
	Date     time.Time `json:"date"`
	Creation time.Time `json:"creation"`
}

type ExpenseRepository interface {
	FindAll() ([]Expense, error)
	Save(expense Expense) (Expense, error)
	FindById(id uint) (Expense, error)
	DeleteExpense(id uint) (int64, error)
}
