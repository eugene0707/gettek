package models

import(
	"time"
)

type Driver struct {
	ID        int `json:"id"`
	Name string `gorm:"not null" json:"name"`
	LicenseNumber string `gorm:"not null" json:"license_number"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

