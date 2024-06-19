package domain

import "time"

type Users struct {
	Id                     int       `gorm:"primaryKey"`
	Name                   string    `gorm:"size:255;not null"`
	Username               string    `gorm:"size:255;not null"`
	Email                  string    `gorm:"size:255;not null;unique"`
	Password               string    `gorm:"not null"`
	Created_at, Updated_at time.Time `gorm:"not null"`
}
