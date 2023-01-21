package entities

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}
