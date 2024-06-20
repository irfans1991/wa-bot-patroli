package domain

import "time"

type Mutasi_masuks struct {
	Id                     int `gorm:"primaryKey"`
	SecurityName           string
	TypeMutasi             string
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
