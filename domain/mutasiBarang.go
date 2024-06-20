package domain

import "time"

type Mutasi_masuks struct {
	Id                     int    `gorm:"primaryKey"`
	Security               string `gorm:"size:255;not null"`
	Type_Mutasi            string `gorm:"size:255;not null"`
	SupplierName           string
	From                   string
	Supplier               string
	PoliceNumber           string
	TotalItems             string
	Unit                   string
	TravelPermit           string
	Remark                 string
	Nota                   string
	Created_at, Updated_at time.Time `gorm:"not null"`
}
