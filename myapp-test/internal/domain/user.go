package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id uint) (User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}
