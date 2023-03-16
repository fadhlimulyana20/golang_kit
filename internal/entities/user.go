package entities

import (
	"template/internal/entities/base"
	"time"
)

type User struct {
	ID         int       `gorm:"primaryKey" json:"id,omitempty"`
	Name       string    `json:"name,omitempty" validate:"required"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	VerifiedAt time.Time `json:"verified_at"`
	Roles      []Role    `json:"roles" gorm:"many2many:user_roles"`
	base.Timestamp
}
