package models

import (
	"time"
)

// User the user model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// TableName for gorm
func (User) TableName() string {
	return "user"
}
