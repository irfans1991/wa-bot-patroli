package domain

import "time"

type Mutasi_masuks struct {
	Id                     int    `gorm:"primaryKey"`
	Security               string `gorm:"size:255;not null"`
	Type_Mutasi            string `gorm:"size:255;not null"`
	Supplier_Name          string `gorm:"size:255;not null"`
	From                   string
	Supplier               string `gorm:"size:255;not null"`
	PoliceNumber           string
	TotalItems             string
	Unit                   string
	TravelPermit           string
	Remark                 string
	Nota                   string
	Created_at, Updated_at time.Time `gorm:"not null"`
}
