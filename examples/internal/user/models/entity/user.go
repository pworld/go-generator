package entity

import (
	"time"
)

// User represents the user entity
type Company struct {
	ID        int        `json:"id"`
	Fullname  string     `json:"fullname" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Phone     string     `json:"phone" validate:"required"`
	Username  string     `json:"username" validate:"required"`
	Password  string     `json:"-" validate:"required"` // "-" in JSON tag to prevent sending the password hash
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"` // Pointer to allow for null value
}
