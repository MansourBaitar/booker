package model

import (
	"time"
)

type User struct {
	Id        uint16    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	Save(user User) (User, error)
	GetUserByEmail(email string, password string) (User, error)
	// FindById(id uint) (User, error)
}
