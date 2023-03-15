package entities

import "template/internal/entities/base"

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	base.Timestamp
}
