package domain

import "time"

type Documents struct {
	Id                     int       `gorm:"primaryKey"`
	Security               string    `gorm:"size:255;not null"`
	Company                string    `gorm:"size:255;not null"`
	Address                string    `gorm:"size:255;not null"`
	Document_type          string    `gorm:"size:255;not null"`
	Sender                 string    `gorm:"size:255;not null"`
	Receiver               string    `gorm:"size:255;not null"`
	Status                 string    `gorm:"size:255;not null"`
	Created_at, Updated_at time.Time `gorm:"not null"`
}
